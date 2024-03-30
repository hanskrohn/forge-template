package actions

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/tui"
)

func (d deleteTemplateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := d.list.Update(msg)
	return d, cmd
}

func (d deleteTemplateModel) Init() tea.Cmd {
	d.list.Init()
	return nil
}

func (d deleteTemplateModel) View() string {
	return d.list.View("Select template to delete:")
}

func DeleteTemplate(userInputData *UserInputData, isDirAction bool) {
	d := newDeleteTemplateModel(userInputData, isDirAction)
	if d == nil {
		return
	}

	p := tea.NewProgram(*d, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

// +++++++++++++++++++
// +     Helpers	 +
// +++++++++++++++++++

func delete(isDirAction bool, templateName string) {
	var path string
	if isDirAction {
		path = directoryTemplatePath
	}else {
		path = fileTemplatePath
	}
	err := os.RemoveAll(path + string(os.PathSeparator) + templateName)

	if err != nil {
		panic(err)
	}
}

func (d deleteTemplateModel) onListSelect(value string) tea.Cmd {
	delete(d.userInputData.isDirAction, value)
	return tea.Quit
}

func newDeleteTemplateModel(userInputData *UserInputData, isDirAction bool) *deleteTemplateModel {
	if userInputData != nil {
		delete(isDirAction, userInputData.TemplateName)
		return nil
	}

	directoryTemplateFileNames, fileTemplateFileNames := common.GetTemplates()

	u := &UserInputData{}
	var choices []string

	if isDirAction {
		choices = directoryTemplateFileNames
	} else {
		choices = fileTemplateFileNames
	}

	u.isDirAction = isDirAction

	d := deleteTemplateModel{
		list: &tui.List{
			Choices: choices,
			Cursor: 0,
		},
		userInputData: u,
	}

	d.list.OnSelect = d.onListSelect

	return &d
}
