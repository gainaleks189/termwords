package tui

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gainaleks189/termwords/internal/dictionary"
)

// One Dark — popular palette (VS Code, Atom).
var (
	numberStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5c6370"))
	wordStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#abb2bf"))
	ghostStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5c6370"))
	successStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#98c379"))
	activeRowStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#2c323c"))
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#61afef"))
	separatorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5c6370"))
	dimStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5c6370"))
)

type Model struct {
	Words     []dictionary.Word
	Start     int
	End       int
	Cursor    int
	Input     textinput.Model
	Wrong     bool
	Answers   map[int]string
	Language  string
	Daily     int
	Completed bool
}

func New(words []dictionary.Word, start, end int, language string, daily int) Model {
	cursor := start
	if cursor > end {
		cursor = end
	}
	ti := textinput.New()
	ti.Prompt = ""
	ti.Placeholder = ""
	ti.CharLimit = 64
	ti.Width = 40
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#abb2bf"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#abb2bf"))
	ti.Focus() // set focus on the struct we store; Init() runs on a copy so the program's model would never get focus otherwise
	return Model{
		Words:    words,
		Start:    start,
		End:      end,
		Cursor:   cursor,
		Input:    ti,
		Answers:  make(map[int]string),
		Language: language,
		Daily:    daily,
	}
}

func (m Model) Init() tea.Cmd {
	fmt.Print("\033[4 q") // steady underline cursor (works in iTerm, Alacritty, Kitty, macOS Terminal, VSCode, Cursor)
	return tea.Batch(
		tea.SetWindowTitle("termwords"),
		m.Input.Focus(),
	)
}

func (m Model) handleEnter() (tea.Model, tea.Cmd) {
	if len(m.Words) == 0 || m.Cursor < 0 || m.Cursor >= len(m.Words) {
		return m, nil
	}
	word := m.Words[m.Cursor]
	typed := strings.TrimSpace(m.Input.Value())
	if typed != word.Answer {
		m.Wrong = true
		m.Input.SetValue("")
		return m, nil
	}
	m.Answers[m.Cursor] = word.Answer
	m.Wrong = false
	m.Input.SetValue("")
	m.Cursor++
	if m.Cursor > m.End {
		m.Completed = true
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > m.Start {
				m.Cursor--
				m.Input.SetValue("")
				m.Wrong = false
			}
			return m, nil
		case "down", "j":
			if m.Cursor < m.End {
				m.Cursor++
				m.Input.SetValue("")
				m.Wrong = false
			}
			return m, nil
		case "enter", "ctrl+j", "ctrl+m":
			return m.handleEnter()
		}
	}
	var cmd tea.Cmd
	m.Input, cmd = m.Input.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	total := len(m.Words)
	header := headerStyle.Render(fmt.Sprintf("termwords · %s · %d", m.Language, m.Daily))
	sep := separatorStyle.Render(strings.Repeat("─", 60))

	maxLen := 0
	for i := m.Start; i <= m.End && i < total; i++ {
		l := utf8.RuneCountInString(m.Words[i].Prompt)
		if l > maxLen {
			maxLen = l
		}
	}

	var lines []string
	lines = append(lines, header)
	lines = append(lines, sep)

	for i := m.Start; i <= m.End && i < total; i++ {
		w := m.Words[i]
		rcount := utf8.RuneCountInString(w.Prompt)
		padding := maxLen - rcount
		var content string
		if val, ok := m.Answers[i]; ok {
			content = successStyle.Render(val)
		} else if i == m.Cursor {
			if m.Wrong {
				typed := m.Input.Value()
				if typed == "" {
					content = ghostStyle.Render(w.Answer)
				} else if strings.HasPrefix(w.Answer, typed) {
					remaining := w.Answer[len(typed):]
					content = wordStyle.Render(typed) + ghostStyle.Render(remaining)
				} else {
					content = wordStyle.Render(typed) + ghostStyle.Render(w.Answer)
				}
			} else {
				content = m.Input.View()
			}
		}
		line := numberStyle.Render(fmt.Sprintf("%03d", i+1)) + "  " +
			wordStyle.Render(w.Prompt) + strings.Repeat(" ", padding) +
			separatorStyle.Render(" │") + "  " + content
		if i == m.Cursor {
			line = activeRowStyle.Render(line)
		}
		lines = append(lines, line)
	}

	lines = append(lines, sep)
	lines = append(lines, dimStyle.Render("↑/↓ move  Enter validate  q quit"))

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
