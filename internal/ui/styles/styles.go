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
			Foreground(Primary)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(Secondary).
			Italic(true)

	// Message styles
	SuccessStyle = lipgloss.NewStyle().
			Foreground(Success).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(Error).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(Warning).
			Bold(true)

	InfoStyle = lipgloss.NewStyle().
			Foreground(Info)

	// Input styles
	PromptStyle = lipgloss.NewStyle().
			Foreground(Primary).
			Bold(true)

	InputStyle = lipgloss.NewStyle().
			Foreground(Accent).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Primary).
			Padding(0, 1)

	// Box styles
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Primary).
			Padding(0, 1)

	HighlightBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.DoubleBorder()).
				BorderForeground(Accent).
				Padding(0, 1)

	// List styles
	SelectedItemStyle = lipgloss.NewStyle().
				Foreground(Primary).
				Bold(true).
				PaddingLeft(1)

	UnselectedItemStyle = lipgloss.NewStyle().
				Foreground(Muted).
				PaddingLeft(2)

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
