package tui

import (
	"encoding/json"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CreateModel struct {
	AppModel
	nameInput    string
	contentInput string
	tagContent   string
	step         int // 0: name, 1: tags, 2: content
	focused      bool
	err          error
}

func NewCreateModel(base AppModel) CreateModel {
	return CreateModel{
		AppModel: base,
		step:     0,
		focused:  true,
	}
}

func (m CreateModel) Init() tea.Cmd {
	return nil
}

func (m CreateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Quit on Esc
			return m, tea.Quit

		case "enter":
			// Handle Enter based on current step
			if m.step == 0 && m.nameInput != "" {
				m.step = 1
				return m, nil
			}
			if m.step == 1 && m.tagContent != "" {
				m.step = 2
				return m, nil
			}
			if m.step == 2 {
				// Add new line to contentInput (multi-line support)
				m.contentInput += "\n"
				return m, nil
			}
		case "ctrl+s":
			// Save the note after content input
			if m.step == 2 {
				return m.saveNote()
			}

		case "ctrl+c":
			// Quit the app
			return m, tea.Quit

		case "backspace":
			// Handle backspace
			if m.step == 0 && len(m.nameInput) > 0 {
				m.nameInput = m.nameInput[:len(m.nameInput)-1]
			} else if m.step == 1 && len(m.tagContent) > 0 {
				m.tagContent = m.tagContent[:len(m.tagContent)-1]
			} else if m.step == 2 && len(m.contentInput) > 0 {
				m.contentInput = m.contentInput[:len(m.contentInput)-1]
			}
		default:
			// Add typed character to the respective input field
			if m.step == 0 {
				m.nameInput += msg.String()
			} else if m.step == 1 {
				m.tagContent += msg.String()
			} else if m.step == 2 {
				m.contentInput += msg.String()
			}
		}
	}

	return m, nil
}

func (m CreateModel) saveNote() (tea.Model, tea.Cmd) {
	// Save note content
	notePath := filepath.Join(m.NotesPath, m.nameInput)
	err := os.WriteFile(notePath, []byte(m.contentInput), 0644)
	if err != nil {
		m.err = err
		return m, nil
	}

	// Create or load metadata
	metadataPath := filepath.Join(m.NotesPath, ".metadata.json")
	metadata := make(map[string]interface{})

	// Check if the metadata file exists
	if _, err := os.Stat(metadataPath); err == nil {
		// If the file exists, try reading it
		fileContent, err := os.ReadFile(metadataPath)
		if err != nil {
			m.err = err
			return m, nil
		}

		// If the file content is empty, initialize it as an empty JSON object
		if len(fileContent) == 0 {
			// Initialize an empty object if it's empty
			fileContent = []byte("{}")
		}

		// Unmarshal the file content into the metadata map
		err = json.Unmarshal(fileContent, &metadata)
		if err != nil {
			m.err = err
			return m, nil
		}
	} else {
		// If the file doesn't exist, initialize it as an empty JSON object
		metadata = make(map[string]interface{})
	}

	// Update metadata with tags for the current note
	metadata[m.nameInput] = map[string]string{
		"tags": m.tagContent,
	}

	// Save the updated metadata back to the file
	metadataBytes, err := json.MarshalIndent(metadata, "", "  ")
	if err != nil {
		m.err = err
		return m, nil
	}

	err = os.WriteFile(metadataPath, metadataBytes, 0644)
	if err != nil {
		m.err = err
		return m, nil
	}

	// Load notes before creating the list model
	notes, err := LoadNotes(m.NotesPath)
	if err != nil {
		m.err = err
		return m, nil
	}

	// Return to the list view after saving
	listModel := NewListModel(m.AppModel, notes)
	return listModel, nil
}

func (m CreateModel) View() string {
	var content string

	// View content based on current step
	if m.step == 0 {
		// Step 0: Enter note name
		content = TitleStyle.Render("Create New Note") + "\n\n"
		content += "Enter note name:\n"
		content += SelectedStyle.Render("> "+m.nameInput+"⎕") + "\n\n"
		content += FooterStyle.Render("Press Enter to continue, Esc to quit")
	} else if m.step == 1 {
		// Step 1: Enter tags
		content = TitleStyle.Render("Create New Note: "+m.nameInput) + "\n\n"
		content += "Enter note tags:\n"
		content += SelectedStyle.Render("> "+m.tagContent+"⎕") + "\n\n"
		content += FooterStyle.Render("Press Enter to continue, Esc to cancel")
	} else if m.step == 2 {
		// Step 2: Enter multi-line content
		content = TitleStyle.Render("Create New Note: "+m.nameInput) + "\n\n"
		content += "Enter note content:\n"
		content += SelectedStyle.Render("> "+m.contentInput+"⎕") + "\n\n"
		content += FooterStyle.Render("Press Enter to add a new line, ctrl+s to save, Esc to cancel")
	}

	// Display any errors
	if m.err != nil {
		content += "\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render("Error: "+m.err.Error())
	}

	return CenterContent(m.Width, m.Height, content)
}
