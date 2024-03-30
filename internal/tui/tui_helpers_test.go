package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
    mock.Mock
}

func (m *Mock) OnConfirm(value string) tea.Cmd {
    m.Called(value)
    return nil
}

func (m *Mock) OnSelect(value string) tea.Cmd {
    m.Called(value)
    return nil
}