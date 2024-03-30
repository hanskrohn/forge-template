package actions

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/create/directories"
	"github.com/hanskrohn/forge-template/internal/tui"
)

func (c createDirectoryModel) Init() tea.Cmd {
	c.list.Init()
	c.textInput.Init(DEFINE_VALUE_FOR_VARIABLE_TEXT)

	return nil
}

func (c createDirectoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if c.templateData.mode == ModeSelectingTemplate{
		cmd = c.list.Update(msg)
	} else {
		cmd = c.textInput.Update(msg)
	}

	return c, cmd
}

func (c createDirectoryModel) View() string {
	if c.templateData.mode == ModeSelectingTemplate {
		return c.list.View("Select template:")
	}

	var variable string
	if c.templateData.variableIndex < len(*c.templateData.variables) {
		variable = (*c.templateData.variables)[c.templateData.variableIndex].VariableName
	}
	
	return c.textInput.View("%s %s",
		DEFINE_VALUE_FOR_VARIABLE_TEXT,
		variable,
	)
}	

func CreateDirectory(userInputData *UserInputData) {
	c := newCreateDirectoryModel(userInputData)
	if c == nil {
		return
	}

	p := tea.NewProgram(*c, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

// +++++++++++++++++++
// +     Helpers	 +
// +++++++++++++++++++

func createDirectory(content string, variables *[]*common.Variable) {
	c := common.ReplaceVariablesInContent(content, variables)
		
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	directories.CreateDirectoryFromTemplate(c, dir)
}

func (c createDirectoryModel) onListSelect(value string) tea.Cmd {
	c.userInputData.TemplateName = value

	content, variables := common.GetFileContentAndVariables(projectTemplatePath + string(os.PathSeparator) + value)

	if len(*variables) == 0 {
		createDirectory(content, variables)
		return tea.Quit
	}

	c.templateData.content = content
	c.templateData.variables = variables

	c.templateData.mode = ModeDefiningVariableNames

	return nil
}

func (c createDirectoryModel) onTextInputConfirm(value string) tea.Cmd {
	(*c.templateData.variables)[c.templateData.variableIndex].VariableValue = value
	c.templateData.variableIndex++

	if c.templateData.variableIndex >= len(*c.templateData.variables) {
		createDirectory(c.templateData.content, c.templateData.variables)
		return tea.Quit
	}

	c.textInput.TextInput.SetValue("")

	return nil
}

func newCreateDirectoryModel(userInputData *UserInputData) *createDirectoryModel {
	mode := ModeSelectingTemplate
	templateData := &templateData{
		variableIndex: 0,
	}

	u := &UserInputData{}
	if userInputData != nil {
		content, variables := common.GetFileContentAndVariables(projectTemplatePath + string(os.PathSeparator) + userInputData.TemplateName)

		if len(*variables) == 0 {
			createDirectory(content, variables)
			return nil
		}

		mode = ModeDefiningContent
		templateData.content = content
		templateData.variables = variables
		u = userInputData
	}
	
	templateData.mode = mode
	projectTemplateFileNames, _ := common.GetTemplates()

	c := createDirectoryModel{
		list: &tui.List{
			Choices: projectTemplateFileNames,
			Cursor: 0,
		},
		textInput: &tui.TextInput{},
		userInputData: u,
		templateData: templateData,
	}

	c.list.OnSelect = c.onListSelect
	c.textInput.OnConfirm = c.onTextInputConfirm

	return &c
}