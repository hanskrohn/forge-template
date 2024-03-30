package actions

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/create/files"
	"github.com/hanskrohn/forge-template/internal/tui"
)

var _, fileTemplatePath = common.GetImportantDirectories()

func (c CreateFileModel) Init() tea.Cmd {
	c.list.Init()
	c.textInput.Init(DEFINE_FILE_NAME_TEXT)

	return nil
}

func (c CreateFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if c.templateData.mode == ModeSelectingTemplate{
		cmd = c.list.Update(msg)
	} else {
		cmd = c.textInput.Update(msg)
	}

	return c, cmd
}

func (c CreateFileModel) View() string {
	if c.templateData.mode == ModeSelectingTemplate {
		return c.list.View("Select template:")
	}

	if c.templateData.mode == ModeDefiningContent {
		var variable string
		if c.templateData.variableIndex < len(*c.templateData.variables) {
			variable = (*c.templateData.variables)[c.templateData.variableIndex].VariableName
		}
		
		return c.textInput.View("%s %s",
			DEFINE_VALUE_FOR_VARIABLE_TEXT,
			variable,
		)
	}

	return c.textInput.View("%s",
		DEFINE_FILE_NAME_TEXT,
	)
}	

func CreateFile(userInputData *UserInputData) {
	c := createFileModel(userInputData)
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

func createFile(content string, variables *[]*common.Variable, fileName string) {
	c := common.ReplaceVariablesInContent(content, variables)
		
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files.CreateFile(c, dir + string(os.PathSeparator) + fileName, true)
}

func (c CreateFileModel) onListSelect(value string) tea.Cmd {
	c.UserInputData.TemplateName = value
	c.UserInputData.FileName = value

	content, variables := common.GetFileContentAndVariables(fileTemplatePath + string(os.PathSeparator) + value)

	c.templateData.content = content
	c.templateData.variables = variables

	c.templateData.mode = ModeDefiningFileName

	return nil
}

func (c CreateFileModel) onTextInputConfirm(value string) tea.Cmd {
	if c.templateData.mode == ModeDefiningFileName {
		if value != "" {
			c.UserInputData.FileName = value
		}

		if len(*c.templateData.variables) == 0 {
			createFile(c.templateData.content, c.templateData.variables, c.UserInputData.FileName)
			return tea.Quit
		}

		c.templateData.mode = ModeDefiningContent
		c.textInput.Init(ENTER_TEMPLATE_VALUE_PLACEHOLDER)

		return nil
	}

	(*c.templateData.variables)[c.templateData.variableIndex].VariableValue = value
	c.templateData.variableIndex++

	if c.templateData.variableIndex >= len(*c.templateData.variables) {
		createFile(c.templateData.content, c.templateData.variables, c.UserInputData.FileName)
		return tea.Quit
	}

	return nil
}

func createFileModel(userInputData *UserInputData) *CreateFileModel {
	mode := ModeSelectingTemplate
	templateData := &templateData{
		variableIndex: 0,
	}

	u := &UserInputData{}
	if userInputData != nil {
		content, variables := common.GetFileContentAndVariables(fileTemplatePath + string(os.PathSeparator) + userInputData.TemplateName)

		if len(*variables) == 0 {
			createFile(content, variables, userInputData.FileName)
			return nil
		}

		mode = ModeDefiningContent
		templateData.content = content
		templateData.variables = variables
		u = userInputData
	}
	
	templateData.mode = mode
	_, fileTemplateName := common.GetTemplates()

	c := CreateFileModel{
		list: &tui.List{
			Choices: fileTemplateName,
			Cursor: 0,
		},
		textInput: &tui.TextInput{},
		UserInputData: u,
		templateData: templateData,
	}

	c.list.OnSelect = c.onListSelect
	c.textInput.OnConfirm = c.onTextInputConfirm

	return &c
}