package files

import (
	"os"
	"path/filepath"
	"strings"
)

func CreateFile(content string, path string, removeEscapeValues bool) {
	if removeEscapeValues{
		placeholder := ""
		content = strings.Replace(content, "\\", placeholder, -1)
		content = strings.Replace(content, "\\", placeholder, -1)
	}

	dir := filepath.Dir(path)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        err := os.MkdirAll(dir, 0755)
        if err != nil {
            panic(err)
        }
    }

	err := os.WriteFile(path, []byte(content), 0660)
	if err != nil {
		panic(err)
	}
}
