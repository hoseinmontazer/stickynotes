package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new note",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter note name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			fmt.Println("Note name cannot be empty")
			return
		}

		fmt.Println("Enter your note (Ctrl+D to save):")
		note, _ := reader.ReadString(0x04)

		notePath := filepath.Join(notesPath, name)
		err := os.WriteFile(notePath, []byte(note), 0644)
		if err != nil {
			fmt.Println("Error saving note:", err)
			return
		}
		fmt.Printf("Note '%s' created successfully!\n", name)
	},
}
