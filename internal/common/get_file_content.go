package common

import (
	"os"
)

func GetFileContent(path string) string {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(contentBytes)
}