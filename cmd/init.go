/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maxbrt/mvnp/internal/ui/textInput"
	"github.com/spf13/cobra"
)

type mvnProject struct {
	GroupId    *textInput.Output
	ArtifactId *textInput.Output
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a new Java project with maven",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		project := mvnProject{
			GroupId:    &textInput.Output{},
			ArtifactId: &textInput.Output{},
		}

		// Create and run the groupId input program
		groupIdModel := textInput.InitialTextInput(project.GroupId, "Enter your GroupId")
		tprogram := tea.NewProgram(groupIdModel)

		// Run the program - the value will be saved to project.GroupId.Output automatically
		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}

		artifactId := textInput.InitialTextInput(project.ArtifactId, "Enter your ArtifactId")
		tprogram = tea.NewProgram(artifactId)

		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}

		if err := generateProject(project); err != nil {
			cobra.CheckErr(err)
		}

		// mvn archetype:generate \
		//     -DgroupId=com.example.helloworld \
		//     -DartifactId=my-first-app \
		//     -DarchetypeArtifactId=maven-archetype-quickstart \
		//     -DarchetypeVersion=1.4 \
		//     -DinteractiveMode=false
		//
	},
}

func generateProject(project mvnProject) error {
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

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
