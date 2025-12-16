package styles

import "github.com/charmbracelet/lipgloss"

// Color palette - Blue theme
var (
	Primary   = lipgloss.Color("39")  // Bright Blue
	Secondary = lipgloss.Color("33")  // Deep Blue
	Success   = lipgloss.Color("42")  // Green
	Error     = lipgloss.Color("196") // Red
	Warning   = lipgloss.Color("214") // Orange
	Info      = lipgloss.Color("117") // Light Blue
	Muted     = lipgloss.Color("240") // Gray
	Accent    = lipgloss.Color("75")  // Sky Blue
)

// Common styles
var (
	// Title styles
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Primary).
			MarginBottom(1)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(Secondary).
			Italic(true)

	// Message styles
	SuccessStyle = lipgloss.NewStyle().
			Foreground(Success).
			Bold(true).
			Padding(0, 1)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(Error).
			Bold(true).
			Padding(0, 1)

	WarningStyle = lipgloss.NewStyle().
			Foreground(Warning).
			Bold(true).
			Padding(0, 1)

	InfoStyle = lipgloss.NewStyle().
			Foreground(Info).
			Padding(0, 1)

	// Input styles
	PromptStyle = lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true).
			MarginBottom(1)

	InputStyle = lipgloss.NewStyle().
			Foreground(Accent).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Primary).
			Padding(0, 1)

	// Box styles
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Primary).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1)

	HighlightBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.DoubleBorder()).
				BorderForeground(Accent).
				Padding(1, 2).
				MarginTop(1).
				MarginBottom(1)

	// List styles
	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(Primary).
				Bold(true).
				PaddingLeft(2)

	UnselectedItemStyle = lipgloss.NewStyle().
				Foreground(Muted).
				PaddingLeft(4)

	// Spinner style
	SpinnerStyle = lipgloss.NewStyle().
			Foreground(Primary)
)

// Helper functions
func SuccessMessage(msg string) string {
	return SuccessStyle.Render("✓ " + msg)
}

func ErrorMessage(msg string) string {
	return ErrorStyle.Render("✗ " + msg)
}

func WarningMessage(msg string) string {
	return WarningStyle.Render("⚠ " + msg)
}

func InfoMessage(msg string) string {
	return InfoStyle.Render("ℹ " + msg)
}

func Title(text string) string {
	return TitleStyle.Render(text)
}

func Subtitle(text string) string {
	return SubtitleStyle.Render(text)
}

func Box(content string) string {
	return BoxStyle.Render(content)
}

func HighlightBox(content string) string {
	return HighlightBoxStyle.Render(content)
}
