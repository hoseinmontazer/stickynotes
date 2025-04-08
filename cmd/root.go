package cmd

import (
	"fmt"
	"os"
	"stickynotes/internal/config"
	"stickynotes/tui" // Make sure this is correct based on your actual folder structure

	"github.com/spf13/cobra"
)

var notesPath string

var rootCmd = &cobra.Command{
	Use:   "stickynote",
	Short: "StickyNotes - Terminal sticky notes for Linux!",
	Long:  `StickyNotes lets you manage quick notes from the terminal using commands like create, list, delete, edit, and show.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Example usage of tui styles
		fmt.Println(tui.TitleStyle.Render("Welcome to StickyNotes!"))
		cmd.Help() // Still invokes Cobra's built-in help if no subcommands are passed
	},
}

func Execute() {
	// Use config.GetNotesPath to get the notes path and handle errors
	var err error
	notesPath, err = config.GetNotesPath()
	if err != nil {
		fmt.Println(tui.CautionStyle.Render(fmt.Sprintf("Error: %v", err)))
		os.Exit(1)
	}

	// Add your subcommands
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(editCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(tui.CautionStyle.Render(fmt.Sprintf("Error executing command: %v", err)))
		os.Exit(1)
	}
}
