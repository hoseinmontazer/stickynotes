// help.go
package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HelpModel struct {
	AppModel
	Selected int
}

func (m HelpModel) Init() tea.Cmd {
	return nil
}

func NewHelpModel(base AppModel) HelpModel {
	return HelpModel{AppModel: base}
}

func (m HelpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle key input, type assertion to HelpModel
	var cmd tea.Cmd
	var updatedModel tea.Model
	updatedModel, cmd = handleKeys(m, msg)

	// Check for key inputs specifically for HelpModel
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			switch m.Selected {
			case 0:
				return NewCreateModel(m.AppModel), nil
			case 1:
				notes, _ := loadNotes(m.NotesPath)
				return NewListModel(m.AppModel, notes), nil
			case 2:
				return m.promptForNoteName(), nil
			}

		}
	}

	// Return the updated HelpModel
	return updatedModel, cmd
}

// promptForNoteName returns a new model for inputting the name of a note
func (m HelpModel) promptForNoteName() tea.Model {
	// Here, we're prompting the user for a note name
	return NewNoteNameInputModel(m.AppModel)
}

func (m HelpModel) View() string {
	options := []string{
		"Create new note",
		"List all notes",
		"View a note",
		"Edit a note",
		"Delete a note",
	}

	var menu []string
	for i, opt := range options {
		if i == m.Selected {
			menu = append(menu, SelectedStyle.Render("> "+opt))
		} else {
			menu = append(menu, NormalStyle.Render("  "+opt))
		}
	}

	return CenterContent(m.Width, m.Height,
		lipgloss.JoinVertical(lipgloss.Left,
			TitleStyle.Render("Sticky Notes"),
			"",
			strings.Join(menu, "\n"),
			"",
			FooterStyle.Render("↑/↓: Navigate • Enter: Select • q: Quit"),
		),
	)
}
