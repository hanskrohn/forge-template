package common

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetImportantDirectories(t *testing.T) {
    projectTemplatePath, fileTemplatePath := GetImportantDirectories()

    homeDir, err := os.UserHomeDir()
    if err != nil {
        t.Fatal(err)
    }

    expectedProjectTemplatePath := filepath.Join(homeDir, ".forge-template", "project_template")
    expectedFileTemplatePath := filepath.Join(homeDir, ".forge-template", "file_template")

    if projectTemplatePath != expectedProjectTemplatePath {
        t.Errorf("Expected project template path to be %s, got %s", expectedProjectTemplatePath, projectTemplatePath)
    }

    if fileTemplatePath != expectedFileTemplatePath {
        t.Errorf("Expected file template path to be %s, got %s", expectedFileTemplatePath, fileTemplatePath)
    }
}