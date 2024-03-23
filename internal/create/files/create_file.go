package files

import (
	"os"
)

func CreateFile(content string, path string) {
	newContent := content

	err := os.WriteFile(path, []byte(newContent), 0660)
	if err != nil {
		panic(err)
	}
}
