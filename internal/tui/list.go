package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (l *List) Init() tea.Cmd {
    return nil
}

func (l *List) Update(msg tea.Msg) tea.Cmd {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
		case "ctrl+c":
			return tea.Quit
		case "enter":
			return l.OnSelect(l.Choices[l.Cursor])
        case "up":
            if l.Cursor > 0 {
                l.Cursor--
            }
        case "down":
            if l.Cursor < len(l.Choices)-1 {
                l.Cursor++
            }
        }
    }
    return nil
}

func (l *List) View(header string) string {
	var s string
	for i, choice := range l.Choices {
		cursor := " "
		if l.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	body := fmt.Sprintf("%s\n\n%s", header, s)
	return body
}