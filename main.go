package main

import (
	"fmt"
	"os"
	"stickynotes/cmd"
	"stickynotes/internal/config"
	"stickynotes/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Setup the notes directory and metadata file
	notesPath, err := config.GetNotesPath()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Handle CLI arguments
	if len(os.Args) < 2 {
		startTUI(notesPath)
		return
	}
	cmd.Execute()
}

func startTUI(notesPath string) {
	// Initialize with help screen
	baseModel := tui.NewAppModel(notesPath)
	// helpModel := tui.NewHelpModel(baseModel)
	notes, err := tui.LoadNotes(notesPath)
	if err != nil {
		fmt.Println("Error loading notes:", err)
		os.Exit(1)
	}
	listModel := tui.NewListModel(baseModel, notes)
	p := tea.NewProgram(
		listModel,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
