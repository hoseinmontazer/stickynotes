package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"stickynotes/tui"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	notesDir = ".stickynotes"
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

	// Handle CLI arguments
	if len(os.Args) < 2 {
		startTUI(notesPath)
		return
	}

	// Fall back to CLI if arguments are provided
	switch os.Args[1] {
	case "new", "n":
		createNewNote(notesPath)
	case "list", "l":
		listNotes(notesPath)
	case "show", "s":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a note name")
			os.Exit(1)
		}
		showNote(notesPath, os.Args[2])
	case "delete", "d":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a note name")
			os.Exit(1)
		}
		deleteNote(notesPath, os.Args[2])
	case "edit", "e":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a note name")
			os.Exit(1)
		}
		editNote(notesPath, os.Args[2])
	case "help", "h":
		startTUI(notesPath)
	default:
		fmt.Println("Unknown command")
		startTUI(notesPath)
	}
}

func startTUI(notesPath string) {
	// Initialize with help screen
	baseModel := tui.NewAppModel(notesPath)
	helpModel := tui.NewHelpModel(baseModel)

	p := tea.NewProgram(
		helpModel,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}

// Original CLI functions remain unchanged below
// --------------------------------------------------

func createNewNote(notesPath string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Note name cannot be empty")
		return
	}

	fmt.Print("Enter your note (Ctrl+D to save):\n")
	note, err := reader.ReadString(0x04)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error reading note:", err)
		return
	}

	notePath := filepath.Join(notesPath, name)
	if err := os.WriteFile(notePath, []byte(note), 0644); err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Printf("Note '%s' created successfully!\n", name)
}

func listNotes(notesPath string) {
	files, err := os.ReadDir(notesPath)
	if err != nil {
		fmt.Println("Error reading notes directory:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No notes found")
		return
	}

	fmt.Println("Your notes:")
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file.Name())
	}
}

func showNote(notesPath, name string) {
	notePath := filepath.Join(notesPath, name)
	content, err := os.ReadFile(notePath)
	if err != nil {
		fmt.Printf("Note '%s' not found\n", name)
		return
	}
	fmt.Printf("\n=== %s ===\n%s\n\n", name, string(content))
}

func deleteNote(notesPath, name string) {
	notePath := filepath.Join(notesPath, name)
	if err := os.Remove(notePath); err != nil {
		fmt.Printf("Error deleting note '%s': %v\n", name, err)
		return
	}
	fmt.Printf("Note '%s' deleted successfully\n", name)
}

func editNote(notesPath, name string) {
	notePath := filepath.Join(notesPath, name)
	if _, err := os.Stat(notePath); os.IsNotExist(err) {
		fmt.Printf("Note '%s' doesn't exist. Create it first.\n", name)
		return
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "nano"
	}

	cmd := exec.Command(editor, notePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error editing note:", err)
		return
	}
	fmt.Printf("Note '%s' updated successfully\n", name)
}
