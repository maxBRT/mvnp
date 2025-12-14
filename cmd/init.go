package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maxbrt/mvnp/internal/ui/spinner"
	"github.com/maxbrt/mvnp/internal/ui/textInput"
	"github.com/spf13/cobra"
)

type mvnProject struct {
	GroupId    *textInput.Output
	ArtifactId *textInput.Output
}

type projectGenResult struct {
	err    error
	stderr string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new Java project with maven",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		project := mvnProject{
			GroupId:    &textInput.Output{},
			ArtifactId: &textInput.Output{},
		}

		// Create and run the groupId input program
		groupIdModel := textInput.InitialModel(project.GroupId, "Enter your GroupId")
		tprogram := tea.NewProgram(groupIdModel)

		// Run the program - the value will be saved to project.GroupId.Output automatically
		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}

		artifactId := textInput.InitialModel(project.ArtifactId, "Enter your ArtifactId")
		tprogram = tea.NewProgram(artifactId)

		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}

		// Run spinner while generating project
		var result projectGenResult
		spinnerModel := spinner.InitialModel("Generating Maven project...", func() error {
			result = generateProject(project)
			return result.err
		})
		tprogram = tea.NewProgram(spinnerModel)
		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}

		// Display stderr output after spinner is done (even if there's an error)
		if result.stderr != "" {
			fmt.Fprint(os.Stderr, result.stderr)
			fmt.Fprintln(os.Stderr) // Add newline for clarity
		}

		// Check for errors
		if result.err != nil {
			cobra.CheckErr(result.err)
		}

	},
}

func generateProject(project mvnProject) projectGenResult {
	cmdName := "mvn"
	if runtime.GOOS == "windows" {
		cmdName = "mvn.cmd"
	}
	args := []string{
		"archetype:generate",
		fmt.Sprintf("-DgroupId=%s", project.GroupId.Output),
		fmt.Sprintf("-DartifactId=%s", project.ArtifactId.Output),
		"-DarchetypeArtifactId=maven-archetype-quickstart",
		"-DarchetypeVersion=RELEASE",
		"-DinteractiveMode=false",
	}

	cmd := exec.Command(cmdName, args...)

	// Capture stderr to display after spinner is done
	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	return projectGenResult{
		err:    err,
		stderr: stderrBuf.String(),
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
