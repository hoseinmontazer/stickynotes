package cmd

import (
	"fmt"
	"os"

	"stickynotes/tui" // Replace with correct import path

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes",
	Run: func(cmd *cobra.Command, args []string) {
		// Read the notes directory
		files, err := os.ReadDir(notesPath)
		if err != nil {
			// Use CautionStyle for error messages
			fmt.Println(tui.CautionStyle.Render(fmt.Sprintf("Error reading notes directory: %v", err)))
			return
		}

		// If there are no notes, display message in SelectedStyle
		if len(files) == 0 {
			fmt.Println(tui.SelectedStyle.Render("No notes found."))
			return
		}

		// Display the header for the list with TitleStyle
		fmt.Println(tui.TitleStyle.Render("Your notes:"))

		// Loop through the files and display them with NormalStyle
		for i, file := range files {
			line := fmt.Sprintf("%d. %s", i+1, file.Name())
			// Render each line with NormalStyle for consistent formatting
			fmt.Println(tui.NormalStyle.Render(line))
		}

		// Footer message with faint styling
		fmt.Println(tui.FooterStyle.Render("\nNotes are stored in ~/.stickynotes/"))
	},
}
