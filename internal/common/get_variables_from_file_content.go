package common

import (
	"regexp"
	"strings"
)

func GetVariablesFromContent(content string) *[]*Variable {
    placeholder := "ESCAPED_SEQUENCE"
    content = strings.Replace(content, "\\<", placeholder, -1)
    content = strings.Replace(content, "\\>", placeholder, -1)
    re := regexp.MustCompile(`<<([^>]*)>>`)

    variableValues := make([]*Variable, 0)
    variablesSeen := make(map[string]bool)

    lines := strings.Split(content, "\n")
    for _, line := range lines {
        matches := re.FindAllStringSubmatch(line, -1)

        for _, match := range matches {
            if _, seen := variablesSeen[match[1]]; !seen {
                variableValue := &Variable{VariableIdentifier: strings.Replace(match[0], placeholder, "\\<<", -1), VariableName: match[1]}
                variableValues = append(variableValues, variableValue)
                variablesSeen[match[1]] = true
            }
        }
    }

    return &variableValues
}