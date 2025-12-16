package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/maxbrt/mvnp/internal/ui/styles"
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
		fmt.Println(styles.InfoMessage("Compiling and running your project..."))
		fmt.Println()
		var c *exec.Cmd
		if len(args) > 0 {
			joinedArgs := strings.Join(args, " ")
			mavenArgs := fmt.Sprintf("-Dexec.args=%s", joinedArgs)
			c = exec.Command("mvn", "compile", "exec:java", mavenArgs)
		} else {
			c = exec.Command("mvn", "compile", "exec:java")
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
}
