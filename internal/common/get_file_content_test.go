package common

import (
	"os"
	"testing"

	"github.com/hanskrohn/forge-template/internal/create/files"
)

func TestGetFileContent(t *testing.T) {
    content := "Hello World"
    fileName := "tempFile.txt"

    files.CreateFile(content, fileName, false)

    output := GetFileContent(fileName)

    if output != content {
        t.Errorf("GetFileContent() = %v, want %v", output, content)
    }

    err := os.RemoveAll(fileName)
    if err != nil {
        t.Fatal(err)
    }
}