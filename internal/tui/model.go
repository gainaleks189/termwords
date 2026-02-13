package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gainaleks189/termwords/internal/dictionary"
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
	Width     int
	Height    int
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
	ti.Focus()
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
	return tea.Batch(
		tea.SetWindowTitle("termwords"),
		m.Input.Focus(),
		func() tea.Msg {
			fmt.Print("\033[4 q") // steady underline cursor for input
			return nil
		},
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

// visibleWindow returns (start, end) of the session indices to render so the cursor
// stays in a fixed position (centered). Uses terminal height; no scrollback — we draw only this slice.
func (m Model) visibleWindow() (start, end int) {
	start, end = m.Start, m.End
	visibleSize := m.Height - 4 // header, sep, sep, hint
	if visibleSize < 3 {
		visibleSize = 3
	}
	sessionSize := m.End - m.Start + 1
	if visibleSize > sessionSize {
		visibleSize = sessionSize
	}
	start = m.Cursor - visibleSize/2
	if start < m.Start {
		start = m.Start
	}
	end = start + visibleSize - 1
	if end > m.End {
		end = m.End
		start = end - visibleSize + 1
		if start < m.Start {
			start = m.Start
		}
	}
	return start, end
}
