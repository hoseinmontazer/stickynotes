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
	Notes         []string
	Selected      int
	TagFilter     string
	Metadata      map[string]map[string]string
	TagSelect     bool
	AvailableTags []string
	err           error
}

func NewListModel(base AppModel, notes []string) ListModel {
	notes = filterOutMetadata(notes)
	metadata, err := loadMetadata(base.NotesPath)
	if err != nil {
		fmt.Println("Error loading metadata:", err)
	}
	return ListModel{
		AppModel: base,
		Notes:    notes,
		Metadata: metadata,
	}
}

func loadMetadata(path string) (map[string]map[string]string, error) {
	metadataPath := filepath.Join(path, ".metadata.json")
	metadata := make(map[string]map[string]string)

	if _, err := os.Stat(metadataPath); err != nil {
		if os.IsNotExist(err) {
			return metadata, nil
		}
		return nil, err
	}

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

	if listModel, ok := updatedModel.(ListModel); ok {
		m = listModel
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.TagSelect {
				// Select a tag
				if len(m.AvailableTags) > 0 && m.Selected < len(m.AvailableTags) {
					selectedTag := m.AvailableTags[m.Selected]
					m.TagFilter = selectedTag
					m.TagSelect = false
					m.Selected = 0
					return m.updateTagFilter([]string{selectedTag}), nil
				}
				return m, nil
			}

			// Select a note
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
			if m.TagSelect {
				// Exit tag selection
				m.TagSelect = false
				return m, nil
			}
			return m, tea.Quit

		case "t":
			return m.handleTagFilter(), nil
		}
	}

	return m, cmd
}

func (m ListModel) handleTagFilter() tea.Model {
	tagSet := make(map[string]struct{})
	for _, tagsMap := range m.Metadata {
		tagSet[tagsMap["tags"]] = struct{}{}
	}

	var tags []string
	for tag := range tagSet {
		tags = append(tags, tag)
	}

	m.TagSelect = true
	m.AvailableTags = tags
	m.Selected = 0
	return m
}

func (m ListModel) updateTagFilter(tags []string) tea.Model {
	var filteredNotes []string
	tagSet := make(map[string]bool)
	for _, t := range tags {
		tagSet[t] = true
	}

	for note, tagsMap := range m.Metadata {
		if tagSet[tagsMap["tags"]] {
			filteredNotes = append(filteredNotes, note)
		}
	}

	if len(filteredNotes) == 0 {
		filteredNotes = m.Notes
	}

	m.Notes = filteredNotes
	return m
}

func (m ListModel) View() string {
	if m.TagSelect {
		// Tag selection view
		var tagList []string
		for i, tag := range m.AvailableTags {
			if i == m.Selected {
				tagList = append(tagList, SelectedStyle.Render("> "+tag))
			} else {
				tagList = append(tagList, NormalStyle.Render("  "+tag))
			}
		}

		return CenterContent(m.Width, m.Height,
			lipgloss.JoinVertical(lipgloss.Left,
				TitleStyle.Render("Available Tags"),
				"",
				strings.Join(tagList, "\n"),
				"",
				FooterStyle.Render("↑/↓: Navigate • Enter: Select Tag • q: Cancel"),
			),
		)
	}

	// Notes list view
	var list []string
	for i, note := range m.Notes {
		if i == m.Selected {
			list = append(list, SelectedStyle.Render("> "+note))
		} else {
			list = append(list, NormalStyle.Render("  "+note))
		}
	}

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
			FooterStyle.Render("↑/↓: Navigate • Enter: Select • t: Filter by Tag • h: help • q: Quit"),
		),
	)
}
