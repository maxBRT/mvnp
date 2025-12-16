package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	g "github.com/maxbrt/mvnp/internal/generate"
	"github.com/maxbrt/mvnp/internal/ui/textInput"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "",
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		root, err := g.FindRoot(dir)
		if err != nil {
			fmt.Println(err)
		}

		packages, err := g.ListAllPackages(root)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(packages)
		output := &textInput.Output{}
		classInput := textInput.InitialModel(output, "Enter your Class Name")
		p := tea.NewProgram(classInput)
		if _, err := p.Run(); err != nil {
			cobra.CheckErr(err)
		}

		fmt.Println(output)

	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
