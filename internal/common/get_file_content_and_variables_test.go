package common

import (
	"os"
	"reflect"
	"testing"

	"github.com/hanskrohn/forge-template/internal/create/files"
)

func TestGetFileContentAndVariables(t *testing.T) {
    content := "Hello <<user>>"
    fileName := "tempFile.txt"

    files.CreateFile(content, fileName, false)

    outputContent, outputVariables := GetFileContentAndVariables(fileName)

    if outputContent != content {
        t.Errorf("GetFileContentAndVariables() content = %v, want %v", outputContent, content)
    }

    expectedVariables := &[]*Variable{
        {VariableIdentifier: "<<user>>", VariableName: "user"},
    }

    if !reflect.DeepEqual(outputVariables, expectedVariables) {
        t.Errorf("GetFileContentAndVariables() variables = %v, want %v", outputVariables, expectedVariables)
    }

    err := os.RemoveAll(fileName)
    if err != nil {
        t.Fatal(err)
    }
}