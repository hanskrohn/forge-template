package state

import (
	"testing"
)

var tests = []struct {
	s    string
	action Action
}{
	{"Unknown", Unknown},
	{"Create Directory Template", CreateDirectoryTemplate},
	{"Create File Template", CreateFileTemplate},
	{"Create Directory From Template", CreateDirectoryFromTemplate},
	{"Create File From Template", CreateFileFromTemplate},
	{"Delete Directory Template", DeleteDirectoryTemplate},
	{"Delete File Template", DeleteFileTemplate},
	{"Save To Github", SaveToGithub},
}

func TestAction_String(t *testing.T) {
    for _, tc := range tests {
		if got := tc.action.String(); got != tc.s {
			t.Errorf("Action.String() = %v, want %v", got, tc.action)
		}
    }
}

func TestStringToAction(t *testing.T) {
    for _, tc := range tests {
		if got := StringToAction(tc.s); got != tc.action {
			t.Errorf("StringToAction() = %v, want %v", got, tc.action)
		}
    }
}

func TestNew(t *testing.T) {
    action := CreateDirectoryTemplate
    s := New(action)

    if s.Action != action {
        t.Errorf("Expected action to be %v, got %v", action, s.Action)
    }
}