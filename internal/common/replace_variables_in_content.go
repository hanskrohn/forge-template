package common

import "strings"

func ReplaceVariablesInContent(content string, variables *[]*Variable) string {
	for _, variable := range *variables {
		content = strings.ReplaceAll(content, variable.VariableIdentifier, variable.VariableValue)
	}

	return content
}