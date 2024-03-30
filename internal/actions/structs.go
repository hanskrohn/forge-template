package actions

import (
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/hanskrohn/forge-template/internal/tui"
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

type UserInputData struct {
	FileName     string
	TemplateName string
}

type templateData struct {
	content       string
	variables 	  *[]*common.Variable
	variableIndex int
	mode		  mode
}



type mainMenuModel struct {
	list 	*tui.List
	state   *state.State
}

type deleteModel struct {
	list 	*tui.List
	state   *state.State
}

type createTemplateModel struct {
	textInput 		*tui.TextInput
	textArea		*tui.TextArea
	templateData 	*CreateTemplateData
	state   		*state.State
	errorTracker 	*errorTracker
}

type CreateTemplateData struct {
	TemplateName string
	Mode     mode
}

type CreateFileOrDirFromTemplateModel struct {
	list 			*tui.List
	textInput 		*tui.TextInput
	templateData 	*CreateFileOrDirFromTemplateData
	state   		*state.State
	errorTracker 	*errorTracker
}

type CreateFileOrDirFromTemplateData struct {
	variables 	   *[]*common.Variable 
	variablesIndex int
	fileContent    string
	TemplateName   string
	FileName  	   string
	Mode      	   mode
}

type errorTracker struct {
	errorHappened bool
	err			  error
}