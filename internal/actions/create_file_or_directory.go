package actions

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/create/directories"
	"github.com/hanskrohn/forge-template/internal/create/files"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/hanskrohn/forge-template/internal/tui"
)

func (m CreateFileOrDirFromTemplateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.templateData.Mode == ModeSelectingTemplate {
		cmd = m.list.Update(msg)
	}else if m.templateData.Mode == ModeDefiningVariableNames{
		if len(*m.templateData.variables) == 0 {
			return m, tea.Quit
		}
		cmd = m.textInput.Update(msg)
	}else if m.templateData.Mode == ModeDefiningFileName{
		cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (m CreateFileOrDirFromTemplateModel) Init() tea.Cmd {
	m.list.Init()

	return nil
}

func (m CreateFileOrDirFromTemplateModel) View() string {
	if m.templateData.Mode == ModeSelectingTemplate {
		return m.list.View("Select template:")
	}else if m.templateData.Mode == ModeDefiningVariableNames {
		var variable string
		if m.templateData.variablesIndex < len(*m.templateData.variables) {
			variable = (*m.templateData.variables)[m.templateData.variablesIndex].VariableName
		}
		
		return m.textInput.View("%s %s",
			DEFINE_VALUE_FOR_VARIABLE_TEXT,
			variable,
		)
	}else if m.templateData.Mode == ModeDefiningFileName {
		return m.textInput.View("%s",
			DEFINE_FILE_NAME_TEXT,
		)
	}

	return UNKNOWN_MODE_ERROR_TEXT
}

func setMode(m *CreateFileOrDirFromTemplateModel, mode mode, placeholder string) {
	m.textInput.Init(placeholder)
	m.textInput.TextInput.Focus()
	m.templateData.Mode = mode
}

func (createFileOrDirFromTemplate *CreateFileOrDirFromTemplateModel) createFileOrDirectory() {
	content := common.ReplaceVariablesInContent(createFileOrDirFromTemplate.templateData.fileContent, createFileOrDirFromTemplate.templateData.variables)
		
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if createFileOrDirFromTemplate.state.Action == state.CreateProjectFromTemplate {
		directories.CreateDirectoryFromTemplate(content, dir)
	}else if createFileOrDirFromTemplate.state.Action == state.CreateFileFromTemplate {
		files.CreateFile(content, dir + string(os.PathSeparator) + createFileOrDirFromTemplate.templateData.FileName, true)
	}
}

func (createFileOrDirFromTemplate *CreateFileOrDirFromTemplateModel) OnListSelect(value string) tea.Cmd {
	createFileOrDirFromTemplate.templateData.TemplateName = value
	createFileOrDirFromTemplate.templateData.FileName = value

	projectTemplatePath, fileTemplatePath := common.GetImportantDirectories()

	var path string
	if createFileOrDirFromTemplate.state.Action == state.CreateProjectFromTemplate {
		path = projectTemplatePath
		setMode(createFileOrDirFromTemplate, ModeDefiningVariableNames, ENTER_TEMPLATE_VALUE_PLACEHOLDER)
	} else if createFileOrDirFromTemplate.state.Action == state.CreateFileFromTemplate {
		path = fileTemplatePath
		setMode(createFileOrDirFromTemplate, ModeDefiningFileName, ENTER_FILE_NAME_PLACEHOLDER)
	}

	content, variables := getFileContentAndVariables(path + string(os.PathSeparator) + value)
	createFileOrDirFromTemplate.templateData.fileContent = content
	createFileOrDirFromTemplate.templateData.variables = variables

	if len(*createFileOrDirFromTemplate.templateData.variables) == 0 && createFileOrDirFromTemplate.templateData.Mode != ModeDefiningFileName{
		createFileOrDirFromTemplate.createFileOrDirectory()
		return tea.Quit
	}

	return nil
}

func getFileContentAndVariables(path string) (string, *[]*common.Variable) {
	content := common.GetFileContent(path)
	variables := common.GetVariablesFromContent(content)

	return content, variables
}

func (createFileOrDirFromTemplate *CreateFileOrDirFromTemplateModel) OnTextInputConfirm(value string) tea.Cmd {
	if createFileOrDirFromTemplate.templateData.Mode == ModeDefiningFileName {
		if value != "" {
			createFileOrDirFromTemplate.templateData.FileName = value
		}

		if len(*createFileOrDirFromTemplate.templateData.variables) == 0 {
			createFileOrDirFromTemplate.createFileOrDirectory()
			return tea.Quit
		}

		setMode(createFileOrDirFromTemplate, ModeDefiningVariableNames, ENTER_TEMPLATE_VALUE_PLACEHOLDER)
		return nil
	}else{
		(*createFileOrDirFromTemplate.templateData.variables)[createFileOrDirFromTemplate.templateData.variablesIndex].VariableValue = value
		createFileOrDirFromTemplate.templateData.variablesIndex++
	
		if  createFileOrDirFromTemplate.templateData.variablesIndex >= len(*createFileOrDirFromTemplate.templateData.variables) {
			createFileOrDirFromTemplate.createFileOrDirectory()
			return tea.Quit
		}
	
		createFileOrDirFromTemplate.textInput.TextInput.SetValue("")
	
		return nil
	}
}

func(createFileOrDirFromTemplate *CreateFileOrDirFromTemplateModel) correctFileOrDirTemplateData(createFileOrDirFromTemplateData *CreateFileOrDirFromTemplateData) {
	projectTemplatePath, fileTemplatePath := common.GetImportantDirectories()

	var path string
	if createFileOrDirFromTemplate.state.Action == state.CreateProjectFromTemplate {
		path = projectTemplatePath
	} else if createFileOrDirFromTemplate.state.Action == state.CreateFileFromTemplate {
		path = fileTemplatePath
	}

	if createFileOrDirFromTemplateData != nil {
		content, variables := getFileContentAndVariables(path + string(os.PathSeparator) + createFileOrDirFromTemplateData.TemplateName)

		createFileOrDirFromTemplateData.fileContent = content
		createFileOrDirFromTemplateData.variables = variables
		createFileOrDirFromTemplateData.variablesIndex = 0

		createFileOrDirFromTemplate.templateData = createFileOrDirFromTemplateData

		if len(*variables) == 0{
			createFileOrDirFromTemplate.createFileOrDirectory()
			return
		}

		setMode(createFileOrDirFromTemplate, createFileOrDirFromTemplateData.Mode, ENTER_TEMPLATE_VALUE_PLACEHOLDER)
	}
}

func CreateFileOrDirectory(s *state.State, createFileOrDirFromTemplateData *CreateFileOrDirFromTemplateData) {
	projectTemplateFileNames, fileTemplateFileNames := common.GetTemplates()

	var choices []string
	if (s.Action == state.CreateProjectFromTemplate) {
		choices = projectTemplateFileNames
	}else{
		choices = fileTemplateFileNames
	}

	createFileOrDirFromTemplateModel := CreateFileOrDirFromTemplateModel{
		list: &tui.List{
			Choices: choices,
			Cursor: 0,
		},
		textInput: &tui.TextInput{},
		errorTracker: &errorTracker{
			errorHappened: false,
		},
		templateData: &CreateFileOrDirFromTemplateData{
			Mode: ModeSelectingTemplate,
			variablesIndex: 0,
		},
		state: s,
	}

	createFileOrDirFromTemplateModel.list.OnSelect = func(value string) tea.Cmd {
		return createFileOrDirFromTemplateModel.OnListSelect(value)
	}

	createFileOrDirFromTemplateModel.textInput.OnConfirm = func(value string) tea.Cmd {
		return createFileOrDirFromTemplateModel.OnTextInputConfirm(value)
	}

	createFileOrDirFromTemplateModel.correctFileOrDirTemplateData(createFileOrDirFromTemplateData)

	p := tea.NewProgram(createFileOrDirFromTemplateModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
