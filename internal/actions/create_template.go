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

func (m createTemplateModel) Init() tea.Cmd {
	m.textInput.Init("Enter Name of template...")
	m.textInput.TextInput.Focus()

	m.textArea.Init("Enter Content...")
	m.textArea.TextArea.Focus()

	return nil
}

func (m createTemplateModel) View() string {
	var errorMessage string
	var format string = "%s\n%s"

	if m.errorTracker.errorHappened {
		errorMessage = "An error occurred: " + m.errorTracker.err.Error()
		m.errorTracker.errorHappened = false
		format = "%s\n\n%s"
	}

	if m.templateData.Mode == ModeDefiningContent {
		return m.textArea.View(format,
			DEFINE_CONTENT_TEXT,
			errorMessage,
		)
	}else if m.templateData.Mode == ModeDefiningName {
		return m.textInput.View(format,
			DEFINE_TEMPLATE_NAME_TEXT,
			errorMessage,
		)
	}

	return UNKNOWN_MODE_ERROR_TEXT
}

func (m createTemplateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.templateData.Mode == ModeDefiningContent {
		cmd = m.textArea.Update(msg)
	}else if m.templateData.Mode == ModeDefiningName {
		cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (createTemplate *createTemplateModel) TextInputOnConfirmFunction(value string) tea.Cmd {
	createTemplate.templateData.TemplateName = value
	createTemplate.templateData.Mode = ModeDefiningContent

	createTemplate.textInput.TextInput.Blur()

	return nil
}

func (createTemplate *createTemplateModel) TextAreaOnConfirmFunction(value string) tea.Cmd {
	projectTemplatePath, fileTemplatePath := common.GetImportantDirectories()
	
	var path string

	if createTemplate.state.Action == state.CreateProjectTemplate {
		_, err := directories.BuildTree(value)
		if err != nil {
			createTemplate.errorTracker.errorHappened = true
			createTemplate.errorTracker.err = err // Let user retry
			return nil
		}

		path = projectTemplatePath + string(os.PathSeparator) + createTemplate.templateData.TemplateName
	}else if createTemplate.state.Action == state.CreateFileTemplate {
		path = fileTemplatePath + string(os.PathSeparator) + createTemplate.templateData.TemplateName
	} 

	files.CreateFile(value, path, false)

	return tea.Quit
}

func(createTemplateModel *createTemplateModel) correctTemplateData(templateData *CreateTemplateData) {
	if templateData != nil {
		createTemplateModel.templateData = templateData
	}
}

func CreateTemplate(s *state.State, templateData *CreateTemplateData) {
	createTemplateModel := createTemplateModel{
		textInput: &tui.TextInput{},
		textArea: &tui.TextArea{},
		templateData: &CreateTemplateData{
			Mode: ModeDefiningName,
		},
		errorTracker: &errorTracker{
			errorHappened: false,
		},
		state: s,
	}

	createTemplateModel.textInput.OnConfirm = func(value string) tea.Cmd {
		return createTemplateModel.TextInputOnConfirmFunction(value)
    }

	createTemplateModel.textArea.OnConfirm = func(value string) tea.Cmd {
		return createTemplateModel.TextAreaOnConfirmFunction(value)
    }

	createTemplateModel.correctTemplateData(templateData)

	p := tea.NewProgram(createTemplateModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
