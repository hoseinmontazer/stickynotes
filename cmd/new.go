package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new",
	Aliases: []string{"n"},
	Short:   "Create a new note",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
