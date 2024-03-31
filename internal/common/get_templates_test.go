package common

import (
	"os"
	"testing"
	"time"

	"github.com/hanskrohn/forge-template/internal/create/files"
)

func TestGetTemplates(t *testing.T) {
    projectTemplatePath, fileTemplatePath := GetImportantDirectories()

    for _, tc := range tests {
        time := time.Now().Format("2000-04-17 11:11:11")

        dName := "tempProject" + time + ".txt"
        fName := "tempFile" + time + ".txt"

        tempDirectory := projectTemplatePath + string(os.PathSeparator) + dName
        tempFile := fileTemplatePath + string(os.PathSeparator) + fName

		files.CreateFile(tc.content, tempDirectory, false)
        files.CreateFile(tc.content, tempFile, false)

        directoryTemplateFileNames, fileTemplateFileNames := GetTemplates()

        if !valueInSlice(dName, directoryTemplateFileNames) {
            t.Errorf("Expected to find %s", dName)
        }

        if !valueInSlice(fName, fileTemplateFileNames){
            t.Errorf("Expected to find %s", fName)
        }

        err := os.RemoveAll(tempDirectory)
		if err != nil {
			t.Fatal(err)
		}

        err = os.RemoveAll(tempFile)
		if err != nil {
			t.Fatal(err)
		}
    }
}

func valueInSlice(value string, slice []string) bool {
    for _, item := range slice {
        if item == value {
            return true
        }
    }
    return false
}