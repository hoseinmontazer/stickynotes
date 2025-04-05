package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NoteNameInputModel struct {
	AppModel
	Input    string
	Error    string
	showList bool // New flag to track if we should show list
}

// NewNoteNameInputModel initializes a new NoteNameInputModel.
func NewNoteNameInputModel(base AppModel) NoteNameInputModel {
	return NoteNameInputModel{
		AppModel: base,
		showList: false, // Initialize as false
	}
}

func (m NoteNameInputModel) Init() tea.Cmd {
	return nil
}

func (m NoteNameInputModel) noteExists(name string) (bool, error) {
	notes, err := loadNotes(m.NotesPath)
	if err != nil {
		return false, err
	}
	for _, note := range notes {
		if note == name {
			return true, nil
		}
	}
	return false, nil
}

func (m NoteNameInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.Input == "" {
				if !m.showList { // Only show list if we haven't already
					notes, err := loadNotes(m.NotesPath)
					if err != nil {
						m.Error = fmt.Sprintf("Error loading notes: %v", err)
						return m, nil
					}
					m.showList = true // Set flag to true
					return NewListModel(m.AppModel, notes), nil
				}
				return m, nil // Do nothing if Enter pressed consecutively
			}

			// Reset showList flag when entering non-empty input
			m.showList = false

			exists, err := m.noteExists(m.Input)
			if err != nil {
				m.Error = fmt.Sprintf("Error checking note: %v", err)
				return m, nil
			}

			if !exists {
				notes, err := loadNotes(m.NotesPath)
				if err != nil {
					m.Error = fmt.Sprintf("Error loading notes: %v", err)
					return m, nil
				}
				m.Error = fmt.Sprintf("Note '%s' not found", m.Input)
				return NewListModel(m.AppModel, notes), nil
			}

			content, err := readNote(m.NotesPath, m.Input)
			if err != nil {
				m.Error = fmt.Sprintf("Error reading note: %v", err)
				return m, nil
			}
			return NewViewModel(m.AppModel, m.Input, content), nil

		case "q", "esc":
			return NewHelpModel(m.AppModel), nil

		case "backspace":
			if len(m.Input) > 0 {
				m.Input = m.Input[:len(m.Input)-1]
			}
			m.Error = ""
			m.showList = false // Reset flag when editing

		default:
			if len(msg.String()) == 1 {
				m.Input += msg.String()
				m.Error = ""
				m.showList = false // Reset flag when typing
			}
		}
	}

	return m, nil
}

func (m NoteNameInputModel) View() string {
	var errorMsg string
	if m.Error != "" {
		errorMsg = "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(m.Error)
	}

	return CenterContent(m.Width, m.Height,
		lipgloss.JoinVertical(lipgloss.Left,
			TitleStyle.Render("Enter Note Name to View:"),
			"",
			m.Input,
			errorMsg,
			"",
			FooterStyle.Render("Press 'Enter' to load note • 'q' to go back • 'Enter' with empty input to see all notes"),
		),
	)
}
