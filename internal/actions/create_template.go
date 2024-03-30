package actions

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hanskrohn/forge-template/internal/create/directories"
	"github.com/hanskrohn/forge-template/internal/create/files"
	"github.com/hanskrohn/forge-template/internal/tui"
)

func (c createTemplateModel) Init() tea.Cmd {
	c.textInput.Init(DEFINE_FILE_NAME_TEXT)
	c.textArea.Init(ENTER_FILE_CONTENT_PLACEHOLDER)

	return nil
}

func (c createTemplateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if c.templateData.mode == ModeDefiningContent {
		cmd = c.textArea.Update(msg)
	}else if c.templateData.mode == ModeDefiningFileName {
		cmd = c.textInput.Update(msg)
	}

	return c, cmd
}

func (c createTemplateModel) View() string {
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

func CreateTemplate(userInputData *UserInputData, isDirAction bool) {
	c := newCreateTemplateModel(userInputData, isDirAction)
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

func (c *createTemplateModel) onTextInputConfirm(value string) tea.Cmd {
	c.userInputData.TemplateName = value
	c.templateData.mode = ModeDefiningContent

	return nil
}

func (c *createTemplateModel) onTextAreaConfirm(value string) tea.Cmd {
	if c.userInputData.isDirAction {
		_, err := directories.BuildTree(value)
		if err != nil {
			c.errorTracker.errorHappened = true
			c.errorTracker.err = err // Let user retry
			return nil
		}
	
		files.CreateFile(value, directoryTemplatePath + string(os.PathSeparator) + c.userInputData.TemplateName, false)
	} else {
		files.CreateFile(value, fileTemplatePath + string(os.PathSeparator) + c.userInputData.TemplateName, false)
	}

	return tea.Quit
}

func newCreateTemplateModel(userInputData *UserInputData, isDirAction bool) *createTemplateModel {
	mode := ModeDefiningFileName
	templateData := &templateData{
		variableIndex: 0,
	}

	u := &UserInputData{}
	if userInputData != nil {
		mode = ModeDefiningContent
		u = userInputData
	}
	
	templateData.mode = mode
	u.isDirAction = isDirAction

	c := createTemplateModel{
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