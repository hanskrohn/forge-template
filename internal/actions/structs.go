package actions

import (
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/hanskrohn/forge-template/internal/tui"
)

type Interacting int

const (
	Unknown Interacting = iota
	InteractingWithDirectoryTemplate
	InteractingWithFileTemplate
)

type createFileModel struct{
	textInput     *tui.TextInput
	list      	  *tui.List
	userInputData *UserInputData
	templateData  *templateData
}

type createDirectoryModel struct{
	textInput     *tui.TextInput
	list      	  *tui.List
	userInputData *UserInputData
	templateData  *templateData
	errorTracker  *errorTracker
}

type createTemplateModel struct{
	textInput     *tui.TextInput
	textArea      *tui.TextArea
	userInputData *UserInputData
	templateData  *templateData
	errorTracker  *errorTracker
}

type deleteTemplateModel struct {
	list 	*tui.List
	userInputData *UserInputData
}

type UserInputData struct {
	FileName     string
	TemplateName string
	isDirAction  bool
}

type templateData struct {
	content       string
	variables 	  *[]*common.Variable
	variableIndex int
	mode		  mode
}

type errorTracker struct {
	errorHappened bool
	err			  error
}



type mainMenuModel struct {
	list 	*tui.List
	state   *state.State
}