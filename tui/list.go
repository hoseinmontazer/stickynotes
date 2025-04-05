package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/lipgloss"
)

type ListModel struct {
	AppModel
	Notes    []string
	Selected int
}

func NewListModel(base AppModel, notes []string) ListModel {
	return ListModel{AppModel: base, Notes: notes}
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Use handleKeys and type-assert to ListModel
	var cmd tea.Cmd
	var updatedModel tea.Model
	updatedModel, cmd = handleKeys(m, msg)

	// Type assertion: Ensure it's a ListModel after handleKeys
	if listModel, ok := updatedModel.(ListModel); ok {
		// Now we have a ListModel, and we can safely work with it
		m = listModel
	} else {
		// If it's not a ListModel, return the unmodified model
		return m, cmd
	}

	// Handle the "enter" key
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" && len(m.Notes) > 0 {
			content, _ := readNote(m.NotesPath, m.Notes[m.Selected])
			// Create and return a new ViewModel to view the selected note
			return NewViewModel(m.AppModel, m.Notes[m.Selected], content), nil
		}
	}

	return m, cmd
}

func (m ListModel) View() string {
	var list []string
	for i, note := range m.Notes {
		if i == m.Selected {
			list = append(list, SelectedStyle.Render("> "+note))
		} else {
			list = append(list, NormalStyle.Render("  "+note))
		}
	}

	return CenterContent(m.Width, m.Height,
		lipgloss.JoinVertical(lipgloss.Left,
			TitleStyle.Render("Your Notes"),
			"",
			strings.Join(list, "\n"),
			"",
			FooterStyle.Render("Enter: View • q: Back"),
		),
	)
}
