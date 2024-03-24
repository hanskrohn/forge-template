package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
    Choices  []string
    Cursor   int
    OnSelect func(string) tea.Cmd
}

type TextInput struct {
    TextInput textinput.Model
    OnConfirm func(string) tea.Cmd
}

type TextArea struct {
    TextArea  textarea.Model
    OnConfirm func(string) tea.Cmd
}
