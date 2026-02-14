package tui

import (
	"fmt"
	"strings"
	"unicode/utf8"

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

func (m Model) View() string {
	if m.Completed {
		return headerStyle.Render("Session complete.") + "\n\n" +
			dimStyle.Render("Press any key to exit.")
	}
	total := len(m.Words)
	header := headerStyle.Render(fmt.Sprintf("termwords · %s · %d", m.Language, m.Daily))
	sep := separatorStyle.Render(strings.Repeat("─", 60))

	visibleStart, visibleEnd := m.visibleWindow()
	if m.Height <= 0 {
		visibleStart, visibleEnd = m.Start, m.End
	}

	maxLen := 0
	for i := visibleStart; i <= visibleEnd && i < total; i++ {
		l := utf8.RuneCountInString(m.Words[i].Prompt)
		if l > maxLen {
			maxLen = l
		}
	}

	var lines []string
	lines = append(lines, header)
	lines = append(lines, sep)

	for i := visibleStart; i <= visibleEnd && i < total; i++ {
		w := m.Words[i]
		rcount := utf8.RuneCountInString(w.Prompt)
		padding := maxLen - rcount
		content := m.renderRowContent(i, w)
		line := numberStyle.Render(fmt.Sprintf("%03d", i+1)) + "  " +
			wordStyle.Render(w.Prompt) + strings.Repeat(" ", padding) +
			separatorStyle.Render(" │") + "  " + content
		if i == m.Cursor {
			line = activeRowStyle.Render(line)
		}
		lines = append(lines, line)
	}

	lines = append(lines, sep)
	lines = append(lines, dimStyle.Render("↑↓ move  Enter validate  q quit"))

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m Model) renderRowContent(i int, w dictionary.Word) string {
	if val, ok := m.Answers[i]; ok {
		return successStyle.Render(val)
	}
	if i != m.Cursor {
		return ""
	}
	if m.Wrong {
		typed := m.Input.Value()
		if typed == "" {
			return ghostStyle.Render(w.Answer)
		}
		// Guard: avoid panic when typed is longer than answer (extra key, paste, etc.)
		if len(typed) <= len(w.Answer) && strings.HasPrefix(w.Answer, typed) {
			remaining := w.Answer[len(typed):]
			return wordStyle.Render(typed) + ghostStyle.Render(remaining)
		}
		return wordStyle.Render(typed) + ghostStyle.Render(w.Answer)
	}
	return m.Input.View()
}
