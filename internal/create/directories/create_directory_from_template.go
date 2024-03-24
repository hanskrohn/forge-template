package directories

import (
	"os"
	"path/filepath"
)

func CreateDirectoryFromTemplate(content string, path string) {
	node, err := BuildTree(content)

	if err != nil {
		panic(err)
	}

	err = createDirectoryHelper(node, path)

	if err != nil {
		panic(err)
	}
}

func createDirectoryHelper(node *Node, path string) error {
	newPath := filepath.Join(path, node.Value)

	if node.isDir {
		if err := os.MkdirAll(newPath, 0755); err != nil {
			return err
		}
	} else {
		file, err := os.Create(newPath)
		if err != nil {
			return err
		}
		
		defer func() {
			if closeErr := file.Close(); closeErr != nil {
				panic(closeErr)
			}
		}()
	}

	for _, child := range node.Children {
		if err := createDirectoryHelper(child, newPath); err != nil {
			return err
		}
	}

	return nil
}