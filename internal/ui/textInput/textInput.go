package textInput

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maxbrt/mvnp/internal/ui/styles"
)

type (
	errMsg error
)

type Output struct {
	Output string
}

type model struct {
	textInput textinput.Model
	err       error
	output    *Output
	header    string
}

func InitialModel(output *Output, header string) model {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40
	ti.Prompt = "│ "
	ti.PromptStyle = lipgloss.NewStyle().Foreground(styles.Primary)
	ti.TextStyle = lipgloss.NewStyle().Foreground(styles.Accent)
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(styles.Primary)

	return model{
		textInput: ti,
		output:    output,
		header:    header,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			// Save the input value to the output pointer before quitting
			m.output.Output = m.textInput.Value()
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	headerStyle := lipgloss.NewStyle().
		Foreground(styles.Primary).
		Bold(true).
		MarginLeft(1)

	inputBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(styles.Primary).
		Padding(0, 1).
		MarginLeft(1)

	helpStyle := lipgloss.NewStyle().
		Foreground(styles.Muted).
		Italic(true).
		MarginLeft(1)

	return fmt.Sprintf(
		"%s\n%s\n%s",
		headerStyle.Render(m.header),
		inputBoxStyle.Render(m.textInput.View()),
		helpStyle.Render("Press Enter to confirm • Esc to cancel"),
	)
}
