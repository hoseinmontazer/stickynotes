# 📝 StickyNotes

**StickyNotes** is a simple, fast, terminal-based sticky notes manager written in Go. It provides both a clean TUI (text user interface) and a minimal CLI to quickly create, view, and delete notes—perfect for quick thoughts, todos, or reminders right from your terminal.

---

## 📦 Features

- ✅ Lightweight and easy to use
- 🧠 Clean TUI built with [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- 📝 CLI support with **Cobra** for easy note management
- 💾 Notes are stored as plain `.txt` files in `~/.stickynotes`
- 🛠 CLI support for scripting or automation
- 🗂 All notes are managed locally — no cloud, no sync
- list notes by tag

---

## 🚀 Installation
#### 🔧 Option 1: Quick Install (Recommended)

Run this one-liner to install StickyNotes automatically:

```bash
curl -sSL https://raw.githubusercontent.com/hoseinmontazer/stickynotes/main/install.sh | bash
```
#### 🔧 Option 1: build and Install manually
```bash
git clone https://github.com/hoseinmontazer/stickynotes.git
cd stickynotes
go build -o stickynotes
mv stickynotes /usr/local/bin
chmod +x /usr/local/bin/stickynotes
```
### Clone and Build

```bash
git clone https://github.com/hoseinmontazer/stickynotes.git
cd stickynotes
go build -o ./build/stickynotes
```

## 📖 Usage

You can run StickyNotes in interactive mode (TUI) or command-line mode (CLI).

### 🖥️ Interactive TUI Mode

Launch the terminal UI:

```bash
stickynotes
```

| Key        | Action                    |
|------------|---------------------------|
| ↑ / ↓      | Navigate menu or notes    |
| Enter      | Select or insert newline  |
| Ctrl+S     | Save note                 |
| Esc        | Cancel/quit input         |
| q          | Go back/quit menu         |
| h          | Return to help screen     |
| t          | List by tag               |

### 🖥️ Command Line Mode

Run the following commands to use StickyNotes via the CLI:

```bash
./stickynotes create         # Create a new note via CLI
./stickynotes list           # List all notes
./stickynotes view <name>    # View a specific note
./stickynotes delete <name>  # Delete a specific note
```


### 📂 Notes Storage
All notes are stored as .txt files in:
```
~/.stickynotes/
```

### 🛠️ Dependencies
This project uses the following libraries:

- [**Bubble Tea**](https://github.com/charmbracelet/bubbletea) – A fun, functional, and stateful way to build terminal UIs.
- [**Lip Gloss**](https://github.com/charmbracelet/lipgloss) – Style definitions for terminal applications.
#### Install via Go:

```bash
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
```
### 🧠 Future Improvements
- Search/filter notes
- Markdown rendering in view mode
- Edit existing notes
- Tagging system

