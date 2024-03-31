package common

import (
	"testing"
)

func TestGetVariablesFromContent(t *testing.T) {
    for _, tc := range tests {
		if output := GetVariablesFromContent(tc.content); !variablesEqualExceptVariableValue(output, tc.variables) {
			t.Errorf("GetVariablesFromContent() = %v, want %v", output, tc.expectedOutput)
		}
    }
}