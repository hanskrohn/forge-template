package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)


func (l *TextArea) Init(placeholder string) tea.Cmd {
	l.TextArea = textarea.New()
	l.TextArea.Placeholder = placeholder
	l.TextArea.SetHeight(10)
	l.TextArea.SetWidth(175)
	l.TextArea.Focus()

    return nil
}

func (l *TextArea) Update(msg tea.Msg) tea.Cmd {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if keyMsg.Type == tea.KeyCtrlC {
			return tea.Quit
		}
		if keyMsg.Type == tea.KeyCtrlS {
			return l.OnConfirm(l.TextArea.Value())
		}
		var cmd tea.Cmd
		l.TextArea, cmd = l.TextArea.Update(msg)
		return cmd
	}
	return nil
}

func (l *TextArea) View(format string, header ...any) string {
	format += "\n\n%s\n\n%s\n%s"
	header = append(header, l.TextArea.View(), "(ctrl+s to save)", "(ctrl+c to quit)")
	return fmt.Sprintf(format, header...)
}