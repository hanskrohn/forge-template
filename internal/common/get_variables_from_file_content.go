package common

import (
	"regexp"
	"strings"
)

func GetVariablesFromContent(content string)  *[]*Variable {
	re := regexp.MustCompile(`<<(.*?)>>`)

    variableValues := make([]*Variable, 0)
    variablesSeen := make(map[string]bool)

    lines := strings.Split(content, "\n")
    for _, line := range lines {
        matches := re.FindAllStringSubmatch(line, -1)

        for _, match := range matches {
            if _, seen := variablesSeen[match[1]]; !seen {
                variableValue := &Variable{VariableIdentifier: match[0], VariableName: match[1]}
                variableValues = append(variableValues, variableValue)
                variablesSeen[match[1]] = true
            }
        }
    }

    return &variableValues
}