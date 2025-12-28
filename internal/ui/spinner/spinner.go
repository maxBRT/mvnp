package spinner

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maxbrt/mvnp/internal/ui/styles"
)

type errMsg error

type completionMsg struct{ err error }

type model struct {
	spinner  spinner.Model
	message  string
	task     func() error
	quitting bool
	err      error
}

func InitialModel(message string, task func() error) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(styles.Primary).Bold(true).MarginLeft(1)
	return model{
		spinner: s,
		message: message,
		task:    task,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		runTask(m.task),
	)
}

func runTask(task func() error) tea.Cmd {
	return func() tea.Msg {
		err := task()
		return completionMsg{err: err}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case completionMsg:
		m.err = msg.err
		m.quitting = true
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return ""
	}

	messageStyle := lipgloss.NewStyle().
		Foreground(styles.Info)

	str := fmt.Sprintf("%s %s", m.spinner.View(), messageStyle.Render(m.message))
	if m.quitting {
		return str
	}
	return str
}
