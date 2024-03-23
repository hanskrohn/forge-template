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
	return m.list.View()
}

func DisplayMainMenu(s *state.State) {
	choices := []string{
		state.CreateProjectTemplate.String(),
		state.CreateFileTemplate.String(),
		state.CreateProjectFromTemplate.String(),
		state.CreateFileFromTemplate.String(),
		state.DeleteProjectTemplate.String(),
		state.DeleteFileTemplate.String(),
		state.SaveToGithub.String(),
	}

	mainMenuModel := mainMenuModel{
		state: s,
		list: &tui.List{
			Choices: choices,
			Cursor: 0,
			OnSelect: func(i int) {
				choice := choices[i]
        		s.Action = state.StringToAction(choice)
			},
		},

	}
	
	p := tea.NewProgram(mainMenuModel)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
