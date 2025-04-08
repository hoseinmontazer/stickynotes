package tui

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor   = lipgloss.Color("#5A56E0")
	SecondaryColor = lipgloss.Color("#FF5F87")
	HighlightColor = lipgloss.Color("#FFD700")
	CautionColor   = lipgloss.Color("#ff0000")

	// Define styles that use the colors
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(HighlightColor).
			MarginBottom(1)

	SomeCommonStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF00"))

	SelectedStyle = lipgloss.NewStyle().
			Foreground(SecondaryColor).
			Bold(true)

	NormalStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	FooterStyle = lipgloss.NewStyle().
			Faint(true).
			MarginTop(1)

	// You can create a specific style for the caution color
	CautionStyle = lipgloss.NewStyle().
			Foreground(CautionColor)
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
