package common

import (
	"testing"
)

func TestReplaceVariablesInContent(t *testing.T) {
    for _, tc := range tests {
		if output := ReplaceVariablesInContent(tc.content, tc.variables); output != tc.expectedOutput {
			t.Errorf("ReplaceVariablesInContent() = %v, want %v", output, tc.expectedOutput)
		}
    }
}