package directories

import (
	"testing"
)

func TestBuildTree(t *testing.T) {
    for _, tc := range tests {
		_, err := BuildTree(tc.input)
		if err != nil && !tc.errorShouldHappen {
			t.Errorf("BuildTree() error = %v, wantErr %v", err, tc.errorShouldHappen)
		}
    }
}