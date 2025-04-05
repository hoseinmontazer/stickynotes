package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ViewModel struct {
	AppModel
	NoteName string
	Content  string
}

func NewViewModel(base AppModel, noteName, content string) ViewModel {
	return ViewModel{AppModel: base, NoteName: noteName, Content: content}
}

func (m ViewModel) Init() tea.Cmd {
	return nil
}

func (m ViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			// Return to the list view
			notes, err := loadNotes(m.NotesPath)
			if err != nil {
				// Handle error (could show an error message)
				return m, nil
			}
			return NewListModel(m.AppModel, notes), nil

		case "e", "edit":
			// Launch the editor for the note
			EditNoteWithEditor(m.NotesPath, m.NoteName)

			// After editing, reload the note and return to the same view with updated content
			content, err := readNote(m.NotesPath, m.NoteName)
			if err != nil {
				// Handle error if note doesn't exist
				m.Content = "Error: Could not load note."
			} else {
				m.Content = content
			}

			// Return the same model with updated content after editing
			return m, nil

		case "d", "delete":
			// Transition to DeleteModel for confirmation
			deleteModel := NewDeleteModel(m.AppModel, m.NoteName)
			return deleteModel, nil

		case "enter":
			// Check if the note content is already loaded
			if m.Content == "" {
				// Load the content of the note with the given NoteName
				content, err := readNote(m.NotesPath, m.NoteName)
				if err != nil {
					// Handle error if note doesn't exist
					m.Content = "Error: Could not load note."
				} else {
					// Display content the first time it's pressed
					m.Content = content
				}
			} else {
				// Show a message that there is no new content or refresh the current content
				m.Content = "No new content. Press again to reload."
			}

			// Optionally, provide feedback after pressing Enter multiple times
			return m, nil
		}
	}

	// Handle other keys through handleKeys if needed
	newModel, cmd := handleKeys(m, msg)
	if updatedModel, ok := newModel.(ViewModel); ok {
		m = updatedModel
	}

	return m, cmd
}

func (m ViewModel) View() string {
	return CenterContent(m.Width, m.Height,
		lipgloss.JoinVertical(lipgloss.Left,
			TitleStyle.Render("View Note: "+m.NoteName),
			"",
			m.Content,
			"",
			FooterStyle.Render("e: Edit • d: Delete • Press 'q' or 'esc' to go back"),
		),
	)
}
