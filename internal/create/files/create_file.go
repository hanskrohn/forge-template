package files

import (
	"os"
)

func CreateFile(content string, path string) {
	err := os.WriteFile(path, []byte(content), 0660)
	if err != nil {
		panic(err)
	}
}
