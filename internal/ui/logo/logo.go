package logo

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maxbrt/mvnp/internal/ui/styles"
)

// 1. Define the Raw ASCII Art (Split into parts for styling)
// We use backticks for multiline strings.
const (
	logoTextMvn = `
  __  __ __     __ _   _ 
 |  \/  |\ \   / /| \ | |
 | |\/| | \ \ / / |  \| |
 | |  | |  \ V /  | |\  |
 |_|  |_|   \_/   |_| \_|`

	logoTextPlus = `   _ 
 _| |_
|_   _|
  |_|`
)

type Model struct {
	quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

// 2. The View Logic using Lipgloss
func (m Model) View() string {
	if m.quitting {
		return ""
	}

	// Style the "MVN" part with blue theme
	mvnStyle := lipgloss.NewStyle().
		Foreground(styles.Primary).
		Bold(true)

	// Style the "+" part with accent blue
	plusStyle := lipgloss.NewStyle().
		Foreground(styles.Accent).
		Bold(true)

	// Join them horizontally
	logo := lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		mvnStyle.Render(logoTextMvn),
		plusStyle.Render(logoTextPlus),
	)

	// Create a container without border, minimal padding
	container := lipgloss.NewStyle().
		Padding(1, 1).
		Render(logo)

	// Return the final view
	return container
}
