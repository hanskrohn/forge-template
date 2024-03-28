package files

import (
	"os"
	"strings"
)

func CreateFile(content string, path string, removeEscapeValues bool) {
	if removeEscapeValues{
		placeholder := ""
		content = strings.Replace(content, "\\", placeholder, -1)
		content = strings.Replace(content, "\\", placeholder, -1)
	}

	err := os.WriteFile(path, []byte(content), 0660)
	if err != nil {
		panic(err)
	}
}
