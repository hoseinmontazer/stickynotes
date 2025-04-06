# 📝 StickyNotes

**StickyNotes** is a simple, fast, terminal-based sticky notes manager written in Go. It provides both a clean TUI (text user interface) and a minimal CLI to quickly create, view, and delete notes—perfect for quick thoughts, todos, or reminders right from your terminal.

---

## 📦 Features

- ✅ Lightweight and easy to use
- 🧠 Clean TUI built with [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- 💾 Notes are stored as plain `.txt` files in `~/.stickynotes`
- 🛠 CLI support for scripting or automation
- 🗂 All notes are managed locally — no cloud, no sync

---

## 🚀 Installation
```bash
git clone https://github.com/hoseinmontazer/stickynotes.git
cd stickynotes
go build -o stickynotes
### Clone and Build

```bash
git clone https://github.com/hoseinmontazer/stickynotes.git
cd stickynotes
go build -o stickynotes

## 📖 Usage

You can run StickyNotes in interactive mode (TUI) or command-line mode (CLI).

### 🖥️ Interactive TUI Mode

Launch the terminal UI:

```bash
stickynotes

🎨 TUI Navigation
Key	Action
↑ / ↓	Navigate menu or notes
Enter	Select or insert newline
Ctrl+S	Save note
Esc	Cancel/quit input
q	Go back/quit menu
h	Return to help screen

### 🖥️ Command Line Mode
./stickynotes create        # Create a new note via CLI
./stickynotes list          # List all notes
./stickynotes view <name>   # View a note
./stickynotes delete <name> # Delete a note


### 📂 Notes Storage
All notes are stored as .txt files in:
~/.stickynotes/


### 🛠️ Dependencies
Bubble Tea
Lip Gloss
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss

### 🧠 Future Improvements
Search/filter notes
Markdown rendering in view mode
Edit existing notes
Tagging system

