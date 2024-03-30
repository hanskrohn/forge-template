package actions

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/hanskrohn/forge-template/internal/tui"
)

func (m deleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := m.list.Update(msg)
	return m, cmd
}

func (m deleteModel) Init() tea.Cmd {
	m.list.Init()
	return nil
}

func (m deleteModel) View() string {
	return m.list.View("Select template to delete:")
}

func (m deleteModel) OnListSelect(value string) tea.Cmd {
	projectTemplatePath, fileTemplatePath := common.GetImportantDirectories()

	var path string
	if m.state.Action == state.DeleteDirectoryTemplate {
		path = projectTemplatePath
	} else if m.state.Action == state.DeleteFileTemplate {
		path = fileTemplatePath
	}

	err := os.RemoveAll(path + string(os.PathSeparator) + value)

	if err != nil {
		panic(err)
	}

	return tea.Quit
}

func DeleteTemplate(s *state.State, fileName string) {
	projectTemplateFileNames, fileTemplateFileNames := common.GetTemplates()

	var choices []string
	if (s.Action == state.CreateDirectoryFromTemplate) {
		choices = projectTemplateFileNames
	}else{
		choices = fileTemplateFileNames
	}

	deleteModel := deleteModel{
		list: &tui.List{
			Choices: choices,
			Cursor: 0,
		},
		state: s,
	}

	deleteModel.list.OnSelect = func(value string) tea.Cmd {
		return deleteModel.OnListSelect(value)
	}

	if fileName != "" {
		deleteModel.OnListSelect(fileName)
		return
	}

	p := tea.NewProgram(deleteModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
