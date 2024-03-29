package actions

import (
	"github.com/hanskrohn/forge-template/internal/common"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/hanskrohn/forge-template/internal/tui"
)

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
	templateData 	*createTemplateData
	state   		*state.State
	errorTracker 	*errorTracker
}

type createTemplateData struct {
	fileName string
	mode     mode
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