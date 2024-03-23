package state

type Action int

const (
	Unknown Action = iota
	CreateProjectTemplate
	CreateFileTemplate
	CreateProjectFromTemplate
	CreateFileFromTemplate
	DeleteProjectTemplate
	DeleteFileTemplate
	SaveToGithub
)

func (a Action) String() string {
    switch a {
    case Unknown:
        return "Unknown"
    case CreateProjectTemplate:
        return "Create Project Template"
    case CreateFileTemplate:
        return "Create File Template"
    case CreateProjectFromTemplate:
        return "Create Project From Template"
    case CreateFileFromTemplate:
        return "Create File From Template"
    case DeleteProjectTemplate:
        return "Delete Project Template"
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
    case "Create Project Template":
        return CreateProjectTemplate
    case "Create File Template":
        return CreateFileTemplate
    case "Create Project From Template":
        return CreateProjectFromTemplate
    case "Create File From Template":
        return CreateFileFromTemplate
    case "Delete Project Template":
        return DeleteProjectTemplate
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

func New() *State {
	return &State{
		Action: Unknown,
	}
}
