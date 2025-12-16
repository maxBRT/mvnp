package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mvnp",
	Short: "Maven Plus - A terminal-friendly Maven workflow for people who prefer CLIs over IDEs",
	Long: `Maven Plus (mvnp) - A terminal-friendly Maven workflow

Maven Plus streamlines your Java development workflow with an intuitive
command-line experience for Maven projects.

Features:
  • Interactive Project Creation - Create new Maven projects with a beautiful TUI
  • Quick Run - Compile and run your Maven project with a single command
  • Quick Test - Run all tests or specific test classes
  • Java Version Selection - Choose Java version during project setup

Common Commands:
  mvnp init         Create a new Maven project
  mvnp run          Compile and run your project
  mvnp test         Run all tests

For more information, visit: https://github.com/maxbrt/mvnp`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Custom initialization for rootCmd if needed
}
