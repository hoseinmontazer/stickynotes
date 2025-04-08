package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [note name]",
	Short: "Show a specific note",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		notePath := filepath.Join(notesPath, args[0])
		content, err := os.ReadFile(notePath)
		if err != nil {
			fmt.Printf("Note '%s' not found\n", args[0])
			return
		}
		fmt.Printf("\n=== %s ===\n%s\n\n", args[0], string(content))
	},
}
