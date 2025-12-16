package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [args...]",
	Short: "Compile and run the current Maven project",
	Long: `Compile and run the current Maven project using 'mvn compile exec:java'.

This command compiles your Java project and executes the main class configured
in your pom.xml (via exec-maven-plugin). Any arguments provided will be passed
to your Java application.

Examples:
  # Run the project without arguments
  mvnp run

  # Run with arguments
  mvnp run arg1 arg2 arg3
`,
	Run: func(cmd *cobra.Command, args []string) {
		joinedArgs := strings.Join(args, " ")

		mavenArgs := fmt.Sprintf("-Dexec.args=%s", joinedArgs)

		c := exec.Command("mvn", "compile", "exec:java", mavenArgs)

		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err := c.Run()
		if err != nil {
			cobra.CheckErr(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
