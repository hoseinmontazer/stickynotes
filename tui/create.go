package tui

import (
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CreateModel struct {
	AppModel
	nameInput    string
	contentInput string
	step         int // 0: name, 1: content
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
			return m, tea.Quit

		case "enter":
			if m.step == 0 && m.nameInput != "" {
				m.step = 1
				return m, nil
			}
			if m.step == 1 {
				return m.saveNote()
			}

		case "ctrl+c":
			return m, tea.Quit

		case "backspace":
			if m.step == 0 {
				if len(m.nameInput) > 0 {
					m.nameInput = m.nameInput[:len(m.nameInput)-1]
				}
			} else {
				if len(m.contentInput) > 0 {
					m.contentInput = m.contentInput[:len(m.contentInput)-1]
				}
			}

		default:
			if m.step == 0 {
				m.nameInput += msg.String()
			} else {
				m.contentInput += msg.String()
			}
		}
	}

	return m, nil
}

func (m CreateModel) saveNote() (tea.Model, tea.Cmd) {
	notePath := filepath.Join(m.NotesPath, m.nameInput)
	err := os.WriteFile(notePath, []byte(m.contentInput), 0644)
	if err != nil {
		m.err = err
		return m, nil
	}

	// Load notes before creating the list model
	notes, err := loadNotes(m.NotesPath)
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

	if m.step == 0 {
		content = TitleStyle.Render("Create New Note") + "\n\n"
		content += "Enter note name:\n"
		content += SelectedStyle.Render("> "+m.nameInput+"⎕") + "\n\n"
		content += FooterStyle.Render("Press Enter to continue, Esc to quit")
	} else {
		content = TitleStyle.Render("Create New Note: "+m.nameInput) + "\n\n"
		content += "Enter note content:\n"
		content += SelectedStyle.Render("> "+m.contentInput+"⎕") + "\n\n"
		content += FooterStyle.Render("Press Enter to save, Esc to cancel")
	}

	if m.err != nil {
		content += "\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render("Error: "+m.err.Error())
	}

	return CenterContent(m.Width, m.Height, content)
}
