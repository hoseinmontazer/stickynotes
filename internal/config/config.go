package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	NotesDir         = ".stickynotes"
	MetadataFileName = ".metadata.json"
)
const (
	DefaultWidth  = 80
	DefaultHeight = 24
)

func GetNotesPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %v", err)
	}

	notesPath := filepath.Join(homeDir, NotesDir)

	// Create notes directory if it doesn't exist
	if err := os.MkdirAll(notesPath, 0755); err != nil {
		return "", fmt.Errorf("error creating notes directory: %v", err)
	}

	metadataFile := filepath.Join(notesPath, MetadataFileName)

	// Create metadata file if it doesn't exist
	if _, err := os.Stat(metadataFile); os.IsNotExist(err) {
		fmt.Println("Metadata file does not exist, creating:", metadataFile)
		if _, err := os.Create(metadataFile); err != nil {
			return "", fmt.Errorf("error creating metadata file: %v", err)
		}
	}

	return notesPath, nil
}
