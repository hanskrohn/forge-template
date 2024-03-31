package tui

import (
	"fmt"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

var choices = []string{"choice1", "choice2", "choice3"}

func TestList_Init(t *testing.T) {
    l := &List{
		Choices: choices,
	}
    cmd := l.Init()

    if cmd != nil {
        t.Errorf("Expected command to be nil, got %v", cmd)
    }
}

func TestList_View(t *testing.T) {
    l  := &List{
		Choices: choices,
	}
    l.Init()

	header := "header"
	var s string
	for i, choice := range l.Choices {
		cursor := " "
		if l.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	expected := fmt.Sprintf("%s\n\n%s", header, s)

    view := l.View(header)

    if view != expected {
        t.Errorf("Expected view to be '%s', got '%s'", expected, view)
    }
}

func TestList_Update(t *testing.T) {
	mockOnSelect := new(Mock)
	l := &List{
		Choices: choices,
		Cursor: 0,
		OnSelect: mockOnSelect.OnSelect,
	}

	msg := tea.KeyMsg{Type: tea.KeyDown}
	cmd := l.Update(msg)
	if cmd != nil {
		t.Errorf("Expected command to be nil for other keys but got %v", cmd)
	}

	msg = tea.KeyMsg{Type: tea.KeyUp}
	cmd = l.Update(msg)
	if cmd != nil {
		t.Errorf("Expected command to be nil for other keys but got %v", cmd)
	}

	msg = tea.KeyMsg{Type: tea.KeyEnter}
	mockOnSelect.On("OnSelect", "choice1").Return(nil)
	l.Update(msg)
	mockOnSelect.AssertCalled(t, "OnSelect", "choice1")

	msg = tea.KeyMsg{Type: tea.KeyCtrlC}
	cmd = l.Update(msg)
	if cmd == nil {
		t.Error("Expected command to be tea.Quit for KeyCtrlC")
	}
}