package actions

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/create/directories"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/hanskrohn/forge-template/internal/tui"
)

func (m mainMenuModel) Init() tea.Cmd {
	directories.CreateRootConfigDirectory()
	return m.list.Init()
}

func (m mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.list.Update(msg)
	return m, cmd
}

func (m mainMenuModel) View() string {
	return m.list.View("Select action:")
}

func DisplayMainMenu(s *state.State) {
	choices := []string{
		state.CreateDirectoryTemplate.String(),
		state.CreateFileTemplate.String(),
		state.CreateDirectoryFromTemplate.String(),
		state.CreateFileFromTemplate.String(),
		state.DeleteDirectoryTemplate.String(),
		state.DeleteFileTemplate.String(),
	}

	mainMenuModel := mainMenuModel{
		state: s,
		list: &tui.List{
			Choices: choices,
			Cursor: 0,
			OnSelect: func(choice string) tea.Cmd {
        		s.Action = state.StringToAction(choice)
				return tea.Quit
			},
		},

	}
	
	p := tea.NewProgram(mainMenuModel)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
