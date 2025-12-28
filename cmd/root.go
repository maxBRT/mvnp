package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mvnp",
	Short: "Maven Plus - A terminal-friendly Maven workflow for people who prefer CLIs over IDEs",
	Long: `Maven Plus (mvnp) - A terminal-friendly Maven workflow

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
