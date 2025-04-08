package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [note name]",
	Short: "Edit a note using the default editor",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notePath := filepath.Join(notesPath, args[0])
		if _, err := os.Stat(notePath); os.IsNotExist(err) {
			fmt.Printf("Note '%s' doesn't exist.\n", args[0])
			return
		}
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}
		c := exec.Command(editor, notePath)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		if err != nil {
			fmt.Println("Error editing note:", err)
		}
	},
}
