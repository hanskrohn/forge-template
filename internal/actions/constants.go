package actions

import "github.com/hanskrohn/forge-template/internal/common"

const DEFINE_TEMPLATE_NAME_TEXT = "Please define a template name. Reference https://github.com/hanskrohn/forge-template for instructions"

const DEFINE_CONTENT_TEXT = "Please define the content of the template. Reference https://github.com/hanskrohn/forge-template for instructions"

const DEFINE_VALUE_FOR_VARIABLE_TEXT = "Define value for variable:"

const DEFINE_FILE_NAME_TEXT = "Define File Name"

const ENTER_TEMPLATE_VALUE_PLACEHOLDER = "Enter Name of template variable..."

const ENTER_FILE_NAME_PLACEHOLDER = "Enter Name of file..."

const ENTER_FILE_CONTENT_PLACEHOLDER = "Enter content..."

const UNKNOWN_MODE_ERROR_TEXT = "UNKNOWN MODE: Please report this issue at https://github.com/hanskrohn/forge-template/issues"

var projectTemplatePath, fileTemplatePath = common.GetImportantDirectories()

type mode int

const (
	ModeDefiningName mode = iota
	ModeDefiningContent
	ModeSelectingTemplate
	ModeDefiningVariableNames
	ModeDefiningFileName
)