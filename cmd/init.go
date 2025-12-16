package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maxbrt/mvnp/internal/ui/multiInput"
	"github.com/maxbrt/mvnp/internal/ui/spinner"
	"github.com/maxbrt/mvnp/internal/ui/textInput"
	projectobjectmodel "github.com/maxbrt/mvnp/internal/xml/projectObjectModel"
	"github.com/spf13/cobra"
)

type mvnProject struct {
	GroupId    *textInput.Output
	ArtifactId *textInput.Output
	Version    *textInput.Output
}

type projectGenResult struct {
	err    error
	stderr string
}

var javaVersions = []list.Item{
	multiInput.Item("8"),
	multiInput.Item("11"),
	multiInput.Item("17"),
	multiInput.Item("21"),
	multiInput.Item("25"),
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new Maven Java project with interactive setup",
	Long: `Create a new Maven Java project with interactive setup.

This command guides you through creating a new Maven project using the
maven-archetype-quickstart. It interactively prompts for:
  - GroupId: Your project's group identifier (e.g., com.example)
  - ArtifactId: Your project's artifact identifier (e.g., my-app)
  - Java Version: Target Java version (8, 11, 17, 21, or 25)

The generated project includes:
  - Standard Maven directory structure
  - Configured pom.xml with selected Java version
  - exec-maven-plugin configured for running with 'mvnp run'
  - Sample App.java with main method
  - Sample unit test

Examples:
  # Create a new project with interactive prompts
  mvnp init

After creation, navigate to your project directory and use:
  mvnp run   - to compile and run your application
  mvnp test  - to run your tests`,
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		project := mvnProject{
			GroupId:    &textInput.Output{},
			ArtifactId: &textInput.Output{},
			Version:    &textInput.Output{},
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

		version := multiInput.InitialModel(javaVersions, "Select your Java Version")
		tprogram = tea.NewProgram(version)

		m, err := tprogram.Run()
		if err != nil {
			cobra.CheckErr(err)
		}
		project.Version.Output = m.(multiInput.Model).Choice

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

		// Read pom.xml
		pom, err := projectobjectmodel.UnmarshalPOM(filepath.Join(project.ArtifactId.Output, "pom.xml"))
		if err != nil {
			cobra.CheckErr(result.err)
		}

		// Set the Java version
		pom.JavaVersionRelease = project.Version.Output
		pom.JavaVersionSource = project.Version.Output

		// Add the exec-maven-plugin
		pom.AddPlugin(projectobjectmodel.Plugin{
			GroupId:    "org.codehaus.mojo",
			ArtifactId: "exec-maven-plugin",
			Version:    "3.6.2",
			Configuration: &projectobjectmodel.PluginConfig{
				ConfigTags: []projectobjectmodel.ConfigTag{
					{
						Name:  "mainClass",
						Value: project.GroupId.Output + ".App",
					},
				},
			},
		})

		// Rewrite pom.xml
		pomBytes, err := pom.MarshalPOM()
		if err != nil {
			cobra.CheckErr(err)
		}

		err = os.WriteFile(filepath.Join(project.ArtifactId.Output, "pom.xml"), pomBytes, 0644)
		if err != nil {
			cobra.CheckErr(err)
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
