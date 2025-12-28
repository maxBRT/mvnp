package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/maxbrt/mvnp/internal/ui/styles"
	"github.com/spf13/cobra"
)

var mainClass string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [args...]",
	Short: "Compile and run the current Maven project",
	Long: `Compile and run the current Maven project using 'mvn compile exec:java'.

This command compiles your Java project and executes the main class configured
in your pom.xml (via exec-maven-plugin). Any arguments provided will be passed
to your Java application.

If exec-maven-plugin is not configured in your pom.xml, you can use the '--main' or '-m' flag to specify the main class.

Examples:
  # Run the project without arguments
  mvnp run

  # Run with arguments
  mvnp run arg1 arg2 arg3
`,
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd

		// Check for --main or -m flag
		if mainClass != "" {
			if len(args) > 0 {
				joinedArgs := strings.Join(args, " ")
				mavenArgs := fmt.Sprintf("-Dexec.args=%s", joinedArgs)
				mainClass = fmt.Sprintf("-Dexec.mainClass=%s", mainClass)
				c = exec.Command("mvn", "compile", "exec:java", mainClass, mavenArgs)
			} else {
				mainClass = fmt.Sprintf("-Dexec.mainClass=%s", mainClass)
				c = exec.Command("mvn", "compile", "exec:java", mainClass)
			}
		} else {
			// Use the default exec-maven-plugin configuration
			if len(args) > 0 {
				joinedArgs := strings.Join(args, " ")
				mavenArgs := fmt.Sprintf("-Dexec.args=%s", joinedArgs)
				c = exec.Command("mvn", "compile", "exec:java", mavenArgs)
			} else {
				c = exec.Command("mvn", "compile", "exec:java")
			}
		}

		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err := c.Run()
		if err != nil {
			fmt.Println()
			fmt.Println(styles.ErrorMessage("Failed to run project"))
			cobra.CheckErr(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&mainClass, "main", "m", "", "The main class to execute (e.g. com.example.App)")
}
