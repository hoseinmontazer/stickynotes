package tui

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DeleteModel struct {
	AppModel
	NoteName     string
	Confirmation string
	Focused      bool
	err          error
}

func NewDeleteModel(base AppModel, noteName string) DeleteModel {
	return DeleteModel{
		AppModel:     base,
		NoteName:     noteName,
		Confirmation: "",
		Focused:      true,
	}
}

func (m DeleteModel) Init() tea.Cmd {
	return nil
}

func (m DeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Return to the list view without deleting
			notes, err := loadNotes(m.NotesPath)
			if err != nil {
				m.err = err
				return m, nil
			}
			return NewListModel(m.AppModel, notes), nil

		case "y": // Confirm deletion
			// Delete the note
			deleteNote(m.NotesPath, m.NoteName)
			// Return to the list view after deletion
			notes, err := loadNotes(m.NotesPath)
			if err != nil {
				m.err = err
				return m, nil
			}
			return NewListModel(m.AppModel, notes), nil

		case "n": // Cancel deletion
			// Return to the list view without deleting the note
			notes, err := loadNotes(m.NotesPath)
			if err != nil {
				m.err = err
				return m, nil
			}
			return NewListModel(m.AppModel, notes), nil
		}
	}

	return m, nil
}

func (m DeleteModel) View() string {
	var content string

	// Use TitleStyle for the header
	content = TitleStyle.Render("Delete Note: "+m.NoteName) + "\n\n"

	// Use NormalStyle for the body content
	content += NormalStyle.Render("Are you sure you want to delete this note? (y/n)\n\n")

	// Use FooterStyle for the footer
	content += FooterStyle.Render("Press 'y' to confirm, 'n' to cancel")

	// If there's an error, display it in a warning style
	if m.err != nil {
		content += "\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render("Error: "+m.err.Error())
	}

	// Center the content and return it
	return CenterContent(m.Width, m.Height, content)
}

func deleteNote(notesPath, name string) {
	notePath := filepath.Join(notesPath, name)
	if err := os.Remove(notePath); err != nil {
		fmt.Printf("Error deleting note '%s': %v\n", name, err)
		return
	}
	fmt.Printf("Note '%s' deleted successfully\n", name)
}
