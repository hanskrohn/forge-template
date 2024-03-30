package actions

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/create/directories"
	"github.com/hanskrohn/forge-template/internal/create/files"
	"github.com/hanskrohn/forge-template/internal/tui"
)

func (c createDirectoryTemplateModel) Init() tea.Cmd {
	c.textInput.Init(DEFINE_FILE_NAME_TEXT)
	c.textArea.Init(ENTER_FILE_CONTENT_PLACEHOLDER)

	return nil
}

func (c createDirectoryTemplateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if c.templateData.mode == ModeDefiningContent {
		cmd = c.textArea.Update(msg)
	}else if c.templateData.mode == ModeDefiningName {
		cmd = c.textInput.Update(msg)
	}

	return c, cmd
}

func (c createDirectoryTemplateModel) View() string {
	var errorMessage string
	var format string = "%s\n%s"

	if c.errorTracker.errorHappened {
		errorMessage = "An error occurred: " + c.errorTracker.err.Error()
		c.errorTracker.errorHappened = false
		format = "%s\n\n%s"
	}

	if c.templateData.mode == ModeDefiningContent {
		return c.textArea.View(format,
			DEFINE_CONTENT_TEXT,
			errorMessage,
		)
	}

	return c.textInput.View(format,
		DEFINE_TEMPLATE_NAME_TEXT,
		errorMessage,
	)
}	

func CreateDirectoryTemplate(userInputData *UserInputData) {
	c := newCreateDirectoryTemplateModel(userInputData)
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

func createTemplate(content string, variables *[]*common.Variable, fileName string) {
	c := common.ReplaceVariablesInContent(content, variables)
		
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files.CreateFile(c, dir + string(os.PathSeparator) + fileName, true)
}

func (c *createDirectoryTemplateModel) onTextInputConfirm(value string) tea.Cmd {
	c.userInputData.TemplateName = value
	c.templateData.mode = ModeDefiningContent

	return nil
}

func (c *createDirectoryTemplateModel) onTextAreaConfirm(value string) tea.Cmd {
	_, err := directories.BuildTree(value)
	if err != nil {
		c.errorTracker.errorHappened = true
		c.errorTracker.err = err // Let user retry
		return nil
	}

	files.CreateFile(value, projectTemplatePath + string(os.PathSeparator) + c.userInputData.TemplateName, false)

	return tea.Quit
}

func newCreateDirectoryTemplateModel(userInputData *UserInputData) *createDirectoryTemplateModel {
	mode := ModeSelectingTemplate
	templateData := &templateData{
		variableIndex: 0,
	}

	u := &UserInputData{}
	if userInputData != nil {
		mode = ModeDefiningContent
		u = userInputData
	}
	
	templateData.mode = mode

	c := createDirectoryTemplateModel{
		textInput: &tui.TextInput{},
		textArea: &tui.TextArea{},
		userInputData: u,
		templateData: templateData,
		errorTracker: &errorTracker{
			errorHappened: false,
		},
	}

	c.textInput.OnConfirm = c.onTextInputConfirm
	c.textArea.OnConfirm = c.onTextAreaConfirm

	return &c
}