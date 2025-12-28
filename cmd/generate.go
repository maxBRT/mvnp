package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	g "github.com/maxbrt/mvnp/internal/generate"
	"github.com/maxbrt/mvnp/internal/ui/multiInput"
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

		types := []list.Item{
			multiInput.Item("Class"),
			multiInput.Item("Interface"),
			multiInput.Item("Abstract Class"),
			multiInput.Item("Enum"),
			multiInput.Item("Record"),
		}

		typeInput := multiInput.InitialModel(types, "Select your Java Type")
		p := tea.NewProgram(typeInput)
		m, err := p.Run()
		if err != nil {
			cobra.CheckErr(err)
		}
		typeChoice := m.(multiInput.Model)

		packages, err := g.ListAllPackages(root)
		if err != nil {
			fmt.Println(err)
		}

		items := make([]list.Item, len(packages))
		for i, p := range packages {
			items[i] = multiInput.Item(p)
		}

		packageInput := multiInput.InitialModel(items, "Select your Java Package")
		p = tea.NewProgram(packageInput)
		m, err = p.Run()
		if err != nil {
			cobra.CheckErr(err)
		}
		packageChoice := m.(multiInput.Model)

		className := &textInput.Output{}
		classInput := textInput.InitialModel(className, "Enter your class name")
		p = tea.NewProgram(classInput)
		if _, err := p.Run(); err != nil {
			cobra.CheckErr(err)
		}

		class := g.NewClassData(strings.TrimPrefix(packageChoice.Choice, "src.main.java."), className.Output, g.ParseType(typeChoice.Choice))
		classTemplate, err := g.ClassTemplate(*class)
		if err != nil {
			fmt.Println(err)
			return
		}

		dir = strings.ReplaceAll(packageChoice.Choice, ".", string(os.PathSeparator))
		path := filepath.Join(dir, class.ClassName+".java")
		if err := os.WriteFile(path, []byte(classTemplate), 0644); err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(genCmd)
}
