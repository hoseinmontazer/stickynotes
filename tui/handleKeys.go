// handleKeys.go
package tui

import tea "github.com/charmbracelet/bubbletea"

// Handle key events for each model
func handleKeys(model tea.Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m := model.(type) {
	case HelpModel:
		// Handling HelpModel specific logic
		if msg, ok := msg.(tea.KeyMsg); ok {
			switch msg.String() {
			case "up", "k":
				if m.Selected > 0 {
					m.Selected--
				}
			case "down", "j":
				if m.Selected < 4 {
					m.Selected++
				}
			case "q":
				return m, tea.Quit
			}
		}
		return m, nil
	case ListModel:
		// Handling ListModel specific logic
		if msg, ok := msg.(tea.KeyMsg); ok {
			switch msg.String() {
			case "up", "k":
				if m.Selected > 0 {
					m.Selected--
				}
			case "down", "j":
				if m.Selected < len(m.Notes)-1 {
					m.Selected++
				}
			case "q":
				return m, tea.Quit
			}
		}
		return m, nil
	case CreateModel:
		// Handling CreateModel specific logic
		// Add relevant key event handling logic here
		return m, nil
	}

	// Return model unmodified if it's not handled
	return model, nil
}
