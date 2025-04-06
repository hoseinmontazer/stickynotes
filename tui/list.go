package tui

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListModel struct {
	AppModel
	Notes     []string
	Selected  int
	TagFilter string
	Metadata  map[string]map[string]string
	err       error
}

func NewListModel(base AppModel, notes []string) ListModel {
	notes = filterOutMetadata(notes)
	metadata, err := loadMetadata(base.NotesPath)
	if err != nil {
		fmt.Println("Error loading metadata:", err)
	}
	return ListModel{AppModel: base, Notes: notes, Metadata: metadata}
}

// Load metadata from .metadata.json
func loadMetadata(path string) (map[string]map[string]string, error) {
	metadataPath := filepath.Join(path, ".metadata.json")
	metadata := make(map[string]map[string]string)

	// Check if metadata file exists
	if _, err := os.Stat(metadataPath); err != nil {
		if os.IsNotExist(err) {
			return metadata, nil
		}
		return nil, err
	}

	// Read and parse the file content
	fileContent, err := os.ReadFile(metadataPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(fileContent, &metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

// Filter out non-note files (e.g., .metadata.json)
func filterOutMetadata(notes []string) []string {
	var filteredNotes []string
	for _, note := range notes {
		if !strings.Contains(note, ".metadata.json") {
			filteredNotes = append(filteredNotes, note)
		}
	}
	return filteredNotes
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var updatedModel tea.Model
	updatedModel, cmd = handleKeys(m, msg)

	// Ensure it's a ListModel after handleKeys
	if listModel, ok := updatedModel.(ListModel); ok {
		m = listModel
	}

	// Handle key actions
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if len(m.Notes) == 0 {
				m.err = fmt.Errorf("no notes available to select")
				return m, nil
			}
			if m.Selected >= 0 && m.Selected < len(m.Notes) {
				content, _ := ReadNote(m.NotesPath, m.Notes[m.Selected])
				return NewViewModel(m.AppModel, m.Notes[m.Selected], content), nil
			}
			m.err = fmt.Errorf("selected index out of range")
			return m, nil

		case "h":
			return NewHelpModel(m.AppModel), nil
		case "q":
			tea.Quit()
		case "t":
			return m.handleTagFilter(), nil
		}
	}

	return m, cmd
}

func (m ListModel) handleTagFilter() tea.Model {
	// Collect unique tags from metadata
	tagSet := make(map[string]struct{})
	for _, tagsMap := range m.Metadata {
		tagSet[tagsMap["tags"]] = struct{}{}
	}

	// Convert tag set to slice
	var tags []string
	for tag := range tagSet {
		tags = append(tags, tag)
	}

	// Update notes based on the selected tag filter
	return m.updateTagFilter(tags)
}

func (m ListModel) updateTagFilter(tags []string) tea.Model {
	var filteredNotes []string

	// Filter notes based on the selected tag
	for _, note := range m.Notes {
		if tagsMap, ok := m.Metadata[note]; ok {
			tag := tagsMap["tags"]
			if m.TagFilter == "" || strings.Contains(tag, m.TagFilter) {
				filteredNotes = append(filteredNotes, note)
			}
		}
	}

	// If no notes match the filter, show all notes
	if len(filteredNotes) == 0 {
		filteredNotes = m.Notes
	}

	m.Notes = filteredNotes
	return m
}

func (m ListModel) View() string {
	var list []string
	for i, note := range m.Notes {
		if i == m.Selected {
			list = append(list, SelectedStyle.Render("> "+note))
		} else {
			list = append(list, NormalStyle.Render("  "+note))
		}
	}

	// Set title based on the tag filter
	var titleText string
	if m.TagFilter != "" {
		titleText = fmt.Sprintf("Your Notes - Tag: %s", m.TagFilter)
	} else {
		titleText = "Your Notes"
	}

	return CenterContent(m.Width, m.Height,
		lipgloss.JoinVertical(lipgloss.Left,
			TitleStyle.Render(titleText),
			"",
			strings.Join(list, "\n"),
			"",
			FooterStyle.Render("↑/↓: Navigate • Enter: Select • t: Filter by Tag • h: help • q: Back"),
		),
	)
}
