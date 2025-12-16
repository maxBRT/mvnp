package multiInput

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maxbrt/mvnp/internal/ui/styles"
)

const listHeight = 14

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(styles.Primary).
			Background(styles.Muted).
			Padding(0, 2).
			MarginLeft(2).
			MarginBottom(1)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color("252"))

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(styles.Primary).
				Bold(true)

	paginationStyle = list.DefaultStyles().
			PaginationStyle.
			PaddingLeft(4).
			Foreground(styles.Muted)

	helpStyle = list.DefaultStyles().
			HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1).
			Foreground(styles.Muted).
			Italic(true)

	quitTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 4).
			Foreground(styles.Success).
			Bold(true)
)

func InitialModel(items []list.Item, title string) Model {
	l := list.New(items, itemDelegate{}, 20, listHeight)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := Model{list: l}
	return m
}

type Item string

func (i Item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	var str string
	if index == m.Index() {
		str = selectedItemStyle.Render("▸ " + string(i))
	} else {
		str = itemStyle.Render("  " + string(i))
	}

	fmt.Fprint(w, str)
}

type Model struct {
	list     list.Model
	Choice   string
	quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.Choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.Choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("✓ Java %s selected", m.Choice))
	}
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View()
}
