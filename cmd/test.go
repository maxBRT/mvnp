package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/maxbrt/mvnp/internal/ui/styles"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test [test-pattern...]",
	Short: "Run tests for the current Maven project",
	Long: `Run tests for the current Maven project using 'mvn test'.

This command executes all tests in your project using Maven's test lifecycle phase.
You can optionally specify test patterns to run specific tests.

Examples:
  # Run all tests
  mvnp test

  # Run specific test class
  mvnp test MyTest

`,
	Run: func(cmd *cobra.Command, args []string) {
		var c *exec.Cmd

		if len(args) > 0 {
			fmt.Println(styles.InfoMessage(fmt.Sprintf("Running tests matching: %s", args[0])))
			// Run specific tests
			testArg := "-Dtest=" + args[0]
			c = exec.Command("mvn", "test", testArg)
		} else {
			fmt.Println(styles.InfoMessage("Running all tests..."))
			// Run all tests
			c = exec.Command("mvn", "test")
		}
		fmt.Println()

		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err := c.Run()
		if err != nil {
			fmt.Println()
			fmt.Println(styles.ErrorMessage("Tests failed"))
			cobra.CheckErr(err)
		}

		fmt.Println()
		fmt.Println(styles.SuccessMessage("All tests passed!"))
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
