package actions

const DEFINE_TEMPLATE_TEXT = `
TODO: Write Text
`

const DEFINE_CONTENT_TEXT = "TODO: Write Text"

const DEFINE_VALUE_FOR_VARIABLE_TEXT = "TODO: Write Text"

const UNKNOWN_MODE_ERROR_TEXT = "UNKNOWN MODE: Please report this issue at TODO"

type mode int

const (
	modeDefiningName mode = iota
	modeDefiningContent
	modeSelectingTemplate
	modeDefiningVariableNames
)