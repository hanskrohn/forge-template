package common

import "os"

func GetImportantDirectories() (string, string) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	projectTemplatePath := homeDir + string(os.PathSeparator) + ".forge-template" + string(os.PathSeparator) + "project_template"
	fileTemplatePath := homeDir + string(os.PathSeparator) + ".forge-template" + string(os.PathSeparator) + "file_template"

	return projectTemplatePath, fileTemplatePath
}