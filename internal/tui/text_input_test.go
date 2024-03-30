package tui

import (
	"fmt"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTextInput_Init(t *testing.T) {
    ti := &TextInput{}
    cmd := ti.Init("Placeholder")

    if cmd != nil {
        t.Errorf("Expected command to be nil, got %v", cmd)
    }

    if ti.TextInput.Placeholder != "Placeholder" {
        t.Errorf("Expected placeholder to be 'Placeholder', got '%s'", ti.TextInput.Placeholder)
    }
}

func TestTextInput_View(t *testing.T) {
    ti := &TextInput{}
    ti.Init("Placeholder")

	baseFormat := "%s"
	expectedFormat := baseFormat + INPUT_FORMAT_BASE_ACTIONS
    header := "header"
    expected := fmt.Sprintf(expectedFormat, header, ti.TextInput.View(), SAVE_ACTION, QUIT_ACTION)

    view := ti.View(baseFormat, header)

    if view != expected {
        t.Errorf("Expected view to be '%s', got '%s'", expected, view)
    }
}

func TestTextInput_Update(t *testing.T) {
	mockOnConfirm := new(Mock)
	ti := &TextInput{OnConfirm: mockOnConfirm.OnConfirm}

	msg := tea.KeyMsg{Runes: []rune("a")}
	cmd := ti.Update(msg)
	if cmd != nil {
		t.Errorf("Expected command to be nil for other keys but got %v", cmd)
	}

	msg = tea.KeyMsg{Type: tea.KeyCtrlS}
	mockOnConfirm.On("OnConfirm", "").Return(nil)
	ti.Update(msg)
	mockOnConfirm.AssertCalled(t, "OnConfirm", "")

	msg = tea.KeyMsg{Type: tea.KeyCtrlC}
	cmd = ti.Update(msg)
	if cmd == nil {
		t.Error("Expected command to be tea.Quit for KeyCtrlC")
	}
}