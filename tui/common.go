package tui

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor   = lipgloss.Color("#5A56E0")
	SecondaryColor = lipgloss.Color("#FF5F87")
	HighlightColor = lipgloss.Color("#FFD700")
	CautionColor   = lipgloss.Color("#ff0000")

	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(HighlightColor).
			MarginBottom(1)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Bold(true)

	NormalStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	FooterStyle = lipgloss.NewStyle().
			Faint(true).
			MarginTop(1)
)

func CenterContent(width, height int, content string) string {
	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
