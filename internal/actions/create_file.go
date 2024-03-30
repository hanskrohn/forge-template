package actions

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/create/files"
	"github.com/hanskrohn/forge-template/internal/tui"
)

func (c createFileModel) Init() tea.Cmd {
	c.list.Init()
	c.textInput.Init(DEFINE_FILE_NAME_TEXT)

	return nil
}

func (c createFileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if c.templateData.mode == ModeSelectingTemplate{
		cmd = c.list.Update(msg)
	} else {
		cmd = c.textInput.Update(msg)
	}

	return c, cmd
}

func (c createFileModel) View() string {
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
	c := newCreateFileModel(userInputData)
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

func (c createFileModel) onListSelect(value string) tea.Cmd {
	c.userInputData.TemplateName = value
	c.userInputData.FileName = value

	content, variables := common.GetFileContentAndVariables(fileTemplatePath + string(os.PathSeparator) + value)

	c.templateData.content = content
	c.templateData.variables = variables

	c.templateData.mode = ModeDefiningFileName

	return nil
}

func (c createFileModel) onTextInputConfirm(value string) tea.Cmd {
	if c.templateData.mode == ModeDefiningFileName {
		if value != "" {
			c.userInputData.FileName = value
		}

		if len(*c.templateData.variables) == 0 {
			createFile(c.templateData.content, c.templateData.variables, c.userInputData.FileName)
			return tea.Quit
		}

		c.templateData.mode = ModeDefiningContent
		c.textInput.Init(ENTER_TEMPLATE_VALUE_PLACEHOLDER)

		return nil
	}

	(*c.templateData.variables)[c.templateData.variableIndex].VariableValue = value
	c.templateData.variableIndex++

	if c.templateData.variableIndex >= len(*c.templateData.variables) {
		createFile(c.templateData.content, c.templateData.variables, c.userInputData.FileName)
		return tea.Quit
	}

	return nil
}

func newCreateFileModel(userInputData *UserInputData) *createFileModel {
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
	_, fileTemplateFileNames := common.GetTemplates()

	c := createFileModel{
		list: &tui.List{
			Choices: fileTemplateFileNames,
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