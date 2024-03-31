package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const INPUT_FORMAT_BASE_ACTIONS = "\n\n%s\n\n%s\n%s"
const SAVE_ACTION = "(ctrl+s to save)"
const QUIT_ACTION = "(ctrl+c to quit)"

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
