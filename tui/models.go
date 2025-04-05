package tui

import (
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

type AppModel struct {
	Width       int
	Height      int
	NotesPath   string
	CurrentView string
	CurrentNote string
}

func NewAppModel(notesPath string) AppModel {
	return AppModel{
		NotesPath:   notesPath,
		CurrentView: "help", // Initial view is the help screen
	}
}

func (m AppModel) Init() tea.Cmd { return nil }

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle key press events
		if msg.String() == "enter" {
			// Handle Enter key press
			notes, err := loadNotes(m.NotesPath)
			if err != nil {
				// Handle error loading notes
				m.CurrentView = "error"
				return m, nil
			}

			// Display the first note for simplicity or a specific one based on selection
			if len(notes) > 0 {
				noteContent, err := readNote(m.NotesPath, notes[0])
				if err != nil {
					// Handle error reading note
					m.CurrentView = "error"
					return m, nil
				}
				m.CurrentNote = noteContent
			}
			// If Enter is pressed consecutively, just reload the current note instead of loading all over again
			m.CurrentView = "note"
		}
	}

	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m.Width = msg.Width
		m.Height = msg.Height
	}
	return m, nil
}

func (m AppModel) View() string {
	switch m.CurrentView {
	case "note":
		return m.CurrentNote
	case "help":
		return "Press Enter to view notes"
	case "error":
		return "Error loading notes"
	default:
		return ""
	}
}

func loadNotes(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var notes []string
	for _, e := range entries {
		notes = append(notes, e.Name())
	}
	return notes, nil
}

func readNote(path, name string) (string, error) {
	content, err := os.ReadFile(filepath.Join(path, name))
	return string(content), err
}
