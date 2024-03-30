package directories

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateDirectoryFromTemplate(t *testing.T) {
    for _, tc := range tests {
		defer func() {
			if r := recover(); r != nil && !tc.errorShouldHappen {
				t.Errorf("CreateDirectoryFromTemplate() panic = %v, wantErr %v", r, tc.errorShouldHappen)
			}
		}()

		CreateDirectoryFromTemplate(tc.input, tc.path)

		node, _ := BuildTree(tc.input)
		checkDirectoryHelper(t, node, tc.path)
    }
}

func checkDirectoryHelper(t *testing.T, node *Node, path string) {
    newPath := filepath.Join(path, node.Value)

    if node.isDir {
        if _, err := os.Stat(newPath); os.IsNotExist(err) {
            t.Errorf("Directory %s was not created", newPath)
        }
    } else {
        if _, err := os.ReadFile(newPath); err != nil {
            t.Errorf("File %s was not created", newPath)
        }
    }

    for _, child := range node.Children {
        checkDirectoryHelper(t, child, newPath)
    }
}