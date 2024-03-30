package tui

import (
	"fmt"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTextArea_Init(t *testing.T) {
    ta := &TextArea{}
    cmd := ta.Init("Placeholder")

    if cmd != nil {
        t.Errorf("Expected command to be nil, got %v", cmd)
    }

    if ta.TextArea.Placeholder != "Placeholder" {
        t.Errorf("Expected placeholder to be 'Placeholder', got '%s'", ta.TextArea.Placeholder)
    }

    if ta.TextArea.Height() != 10 {
        t.Errorf("Expected height to be 10, got %d", ta.TextArea.Height())
    }
}

func TestTextArea_View(t *testing.T) {
    ta := &TextArea{}
    ta.Init("Placeholder")

	baseFormat := "%s"
	expectedFormat := baseFormat + INPUT_FORMAT_BASE_ACTIONS
    header := "header"
    expected := fmt.Sprintf(expectedFormat, header, ta.TextArea.View(), SAVE_ACTION, QUIT_ACTION)

    view := ta.View(baseFormat, header)

    if view != expected {
        t.Errorf("Expected view to be '%s', got '%s'", expected, view)
    }
}

func TestTextArea_Update(t *testing.T) {
	mockOnConfirm := new(Mock)
	ta := &TextArea{OnConfirm: mockOnConfirm.OnConfirm}

	msg := tea.KeyMsg{Runes: []rune("a")}
	cmd := ta.Update(msg)
	if cmd != nil {
		t.Errorf("Expected command to be nil for other keys but got %v", cmd)
	}

	msg = tea.KeyMsg{Type: tea.KeyCtrlS}
	mockOnConfirm.On("OnConfirm", "").Return(nil)
	ta.Update(msg)
	mockOnConfirm.AssertCalled(t, "OnConfirm", "")

	msg = tea.KeyMsg{Type: tea.KeyCtrlC}
	cmd = ta.Update(msg)
	if cmd == nil {
		t.Error("Expected command to be tea.Quit for KeyCtrlC")
	}
}