package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gainaleks189/termwords/internal/debug"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if debug.Log != nil {
		if _, ok := msg.(tea.KeyMsg); ok {
			debug.Log.Println("Update key — Cursor:", m.Cursor, "Start:", m.Start, "End:", m.End, "Input:", m.Input.Value())
		}
	}
	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.Width = size.Width
		m.Height = size.Height
		return m, nil
	}
	if m.Completed {
		if _, ok := msg.(tea.KeyMsg); ok {
			return m, tea.Quit
		}
		return m, nil
	}
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		switch keyMsg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up":
			if m.Cursor > m.Start {
				m.Cursor--
				m.Input.SetValue("")
				m.Wrong = false
			}
			return m, nil
		case "down":
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
