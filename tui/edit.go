package tui

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type EditModel struct {
	AppModel
	NoteName string
	NotePath string
	Content  string
	Focused  bool
	err      error
}

func NewEditModel(base AppModel, noteName string) (EditModel, tea.Cmd) {
	// Build the full path to the note file
	notePath := filepath.Join(base.NotesPath, noteName)
	content, err := ReadNote(base.NotesPath, noteName)
	if err != nil {
		return EditModel{}, tea.Quit
	}

	// Return the initial EditModel with content
	return EditModel{
		AppModel: base,
		NoteName: noteName,
		NotePath: notePath,
		Content:  content,
		Focused:  true,
		err:      err,
	}, nil
}

func (m EditModel) Init() tea.Cmd {
	return nil
}

func (m EditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Return to the list view on Esc
			notes, err := LoadNotes(m.NotesPath)
			if err != nil {
				m.err = err
				return m, nil
			}
			return NewListModel(m.AppModel, notes), nil

		case "e", "edit":
			// Launch the system editor for editing
			EditNoteWithEditor(m.NotesPath, m.NoteName)

			// After editing, reload the note content and return to the same model
			content, err := ReadNote(m.NotesPath, m.NoteName)
			if err != nil {
				m.err = err
				return m, nil
			}

			// Update content after editing
			m.Content = content
			return m, nil

		case "enter":
			// Handle saving the note (optional, you can add save functionality here)
			return m.saveNote()

		case "backspace":
			// Handle backspace if necessary
		}
	}

	return m, nil
}

func (m EditModel) saveNote() (tea.Model, tea.Cmd) {
	// Save the note (use any format you prefer, here it overwrites the current note)
	err := os.WriteFile(m.NotePath, []byte(m.Content), 0644)
	if err != nil {
		m.err = err
		return m, nil
	}

	// Reload notes after saving
	notes, err := LoadNotes(m.NotesPath)
	if err != nil {
		m.err = err
		return m, nil
	}

	// Return to the list view after saving
	return NewListModel(m.AppModel, notes), nil
}

func (m EditModel) View() string {
	var content string

	// Display the content and provide instructions
	content = TitleStyle.Render("Edit Note: "+m.NoteName) + "\n\n"
	content += m.Content + "\n\n"
	content += FooterStyle.Render("Press 'e' to edit, Esc to go back, Enter to save")

	// If there's an error, display it
	if m.err != nil {
		content += "\n\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render("Error: "+m.err.Error())
	}

	return CenterContent(m.Width, m.Height, content)
}

// EditNoteWithEditor launches the external editor
func EditNoteWithEditor(notesPath, name string) {
	// Build the full path to the note file
	notePath := filepath.Join(notesPath, name)

	// Check if the note exists
	if _, err := os.Stat(notePath); os.IsNotExist(err) {
		// If the note doesn't exist, inform the user and return
		fmt.Printf("Note '%s' doesn't exist. Create it first.\n", name)
		return
	}

	// Get the default editor from the environment variable
	editor := os.Getenv("EDITOR")
	if editor == "" {
		// If no editor is set, use "nano" by default
		editor = "nano"
	}

	// Prepare the command to launch the editor
	cmd := exec.Command(editor, notePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the editor
	if err := cmd.Run(); err != nil {
		// If there's an error running the editor, print it
		fmt.Println("Error editing note:", err)
		return
	}

	// If editing was successful, print a confirmation message
	fmt.Printf("Note '%s' updated successfully\n", name)
}
