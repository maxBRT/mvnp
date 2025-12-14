/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
			GroupId: &textInput.Output{},
		}

		// Create and run the groupId input program
		groupIdModel := textInput.InitialTextInput(project.GroupId, "Enter your GroupId")
		tprogram := tea.NewProgram(groupIdModel)

		// Run the program - the value will be saved to project.GroupId.Output automatically
		_, err := tprogram.Run()
		if err != nil {
			cobra.CheckErr(err)
		}

		// Access the user's input from the Output struct
		fmt.Println("GroupId:", project.GroupId.Output)

		// mvn archetype:generate \
		//     -DgroupId=com.example.helloworld \
		//     -DartifactId=my-first-app \
		//     -DarchetypeArtifactId=maven-archetype-quickstart \
		//     -DarchetypeVersion=1.4 \
		//     -DinteractiveMode=false
		//
	},
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
