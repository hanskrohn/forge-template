package common

import (
	"os"
)

func GetTemplates() ([]string, []string) {
	projectTemplatePath, fileTemplatePath := GetImportantDirectories()

	directoryTemplateFileNames := getFileNamesForTemplates(projectTemplatePath)
	fileTemplateFileNames := getFileNamesForTemplates(fileTemplatePath)

	return directoryTemplateFileNames, fileTemplateFileNames
}

func getFileNamesForTemplates(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fileNames := make([]string, len(files))

	for i, files := range files {
		fileNames[i] = files.Name()
	}

	return fileNames
}