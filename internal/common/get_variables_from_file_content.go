package common

import (
	"regexp"
	"strings"
)

func GetVariablesFromContent(content string)  *[]*Variable{
	re := regexp.MustCompile(`<<(.*?)>>`)

	variableValues := make([]*Variable, 0)
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
	
		for _, match := range matches {
			variableValue := &Variable{ VariableIdentifier: match[0]}
			variableValues = append(variableValues, variableValue)
		}
	}

	return &variableValues
}