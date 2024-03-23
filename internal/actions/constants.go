package actions

const DEFINE_TEMPLATE_TEXT = `
TODO: Write Text
`

const DEFINE_CONTENT_TEXT = "TODO: Write Text"

type mode int

const (
	modeDefiningName mode = iota
	modeDefiningContent
)