package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [note name]",
	Short: "Delete a note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notePath := filepath.Join(notesPath, args[0])
		err := os.Remove(notePath)
		if err != nil {
			fmt.Printf("Error deleting note '%s': %v\n", args[0], err)
			return
		}
		fmt.Printf("Note '%s' deleted successfully\n", args[0])
	},
}
