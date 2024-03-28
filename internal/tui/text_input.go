package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)


func (l *TextInput) Init(placeholder string) tea.Cmd {
	l.TextInput = textinput.New()
	l.TextInput.Placeholder = placeholder

    return nil
}

func (l *TextInput) Update(msg tea.Msg) tea.Cmd {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.Type == tea.KeyCtrlC {
			return tea.Quit
		}
		if keyMsg.Type == tea.KeyCtrlS {
			return l.OnConfirm(l.TextInput.Value())
		}
		var cmd tea.Cmd
		l.TextInput, cmd = l.TextInput.Update(msg)
		return cmd
	}
	return nil
}

func (l *TextInput) View(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}