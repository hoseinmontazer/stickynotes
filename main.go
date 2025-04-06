package main

import (
	"fmt"
	"os"
	"path/filepath"
	"stickynotes/cli"
	"stickynotes/tui"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	notesDir         = ".stickynotes"
	metadataFileName = ".metadata.json"
)

func main() {
	// Setup notes directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	notesPath := filepath.Join(homeDir, notesDir)
	if err := os.MkdirAll(notesPath, 0755); err != nil {
		fmt.Println("Error creating notes directory:", err)
		os.Exit(1)
	}
	metadataFile := filepath.Join(notesPath, metadataFileName)
	if _, err := os.Stat(metadataFile); os.IsNotExist(err) {
		fmt.Println("File does not exist, creating:", metadataFile)
		file, err := os.Create(metadataFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer file.Close()
	}

	// Handle CLI arguments
	if len(os.Args) < 2 {
		startTUI(notesPath)
		return
	}
	cli.Start(notesPath)
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
