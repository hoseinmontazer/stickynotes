// cli/commands.go
package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Start(notesPath string) {
	// Check command-line arguments
	if len(os.Args) < 2 {
		displayHelp()
		listNotes(notesPath)
		return
	}

	switch os.Args[1] {
	case "new", "n":
		createNewNotes(notesPath)
	case "list", "l":
		listNotes(notesPath)
	case "show", "s":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a note name")
			return
		}
		showNote(notesPath, os.Args[2])
	case "delete", "d":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a note name")
			return
		}
		deleteNote(notesPath, os.Args[2])
	case "edit", "e":
		if len(os.Args) < 3 {
			fmt.Println("Please specify a note name")
			return
		}
		editNote(notesPath, os.Args[2])
	case "help", "h":
		displayHelp()
	default:
		fmt.Println("Unknown command")
		displayHelp()
	}
}

func displayHelp() {
	fmt.Println(`stickyNotes - Terminal sticky notes for linux!
    Usage:
        stickynote [command]
    Commands:
        new (n)     - Create a new note
        list (l)    - List all notes
        show (s)    - Show a specific note
        delete (d)  - Delete a note
        edit (e)    - Edit a note
        help (h)    - Show this help message
    Notes are stored in ~/.steakynotes/`)
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

func createNewNotes(notesPath string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter note name:")
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
