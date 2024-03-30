package state

type Action int

const (
	Unknown Action = iota
	CreateDirectoryTemplate
	CreateFileTemplate
	CreateDirectoryFromTemplate
	CreateFileFromTemplate
	DeleteDirectoryTemplate
	DeleteFileTemplate
	SaveToGithub
)

func (a Action) String() string {
    switch a {
    case Unknown:
        return "Unknown"
    case CreateDirectoryTemplate:
        return "Create Directory Template"
    case CreateFileTemplate:
        return "Create File Template"
    case CreateDirectoryFromTemplate:
        return "Create Directory From Template"
    case CreateFileFromTemplate:
        return "Create File From Template"
    case DeleteDirectoryTemplate:
        return "Delete Directory Template"
    case DeleteFileTemplate:
        return "Delete File Template"
    case SaveToGithub:
        return "Save To Github"
    default:
        return "Unknown"
    }
}

func StringToAction(s string) Action {
    switch s {
    case "Unknown":
        return Unknown
    case "Create Directory Template":
        return CreateDirectoryTemplate
    case "Create File Template":
        return CreateFileTemplate
    case "Create Directory From Template":
        return CreateDirectoryFromTemplate
    case "Create File From Template":
        return CreateFileFromTemplate
    case "Delete Directory Template":
        return DeleteDirectoryTemplate
    case "Delete File Template":
        return DeleteFileTemplate
    case "Save To Github":
        return SaveToGithub
    default:
        return Unknown
    }
}

type State struct {
	Action Action
}

func New(action Action) *State {
	return &State{
		Action: action,
	}
}
