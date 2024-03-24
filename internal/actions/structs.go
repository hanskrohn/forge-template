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
	templateData 	*createFileOrDirFromTemplateData
	state   		*state.State
	errorTracker 	*errorTracker
}

type createFileOrDirFromTemplateData struct {
	variables 	   *[]*common.Variable 
	variablesIndex int
	fileName  	   string
	fileContent    string
	mode      	   mode
}

type errorTracker struct {
	errorHappened bool
	err			  error
}