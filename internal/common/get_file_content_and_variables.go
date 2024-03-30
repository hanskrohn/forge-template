package common

func GetFileContentAndVariables(path string) (string, *[]*Variable) {
	content := GetFileContent(path)
	variables := GetVariablesFromContent(content)

	return content, variables
}