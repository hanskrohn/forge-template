package common

import "reflect"

var tests = []struct {
	content   	   string
	variables 	   *[]*Variable
	expectedOutput string
}{
	{
		content:  "<<user>> <<job>>.",
		variables: &[]*Variable{
			{VariableIdentifier: "<<user>>", VariableName: "user", VariableValue: "John"},
			{VariableIdentifier: "<<job>>", VariableName: "job", VariableValue: "developer"},
		},
		expectedOutput: "John developer.",
	},

	{
		content:  "<<user>><<job>><<user>>",
		variables: &[]*Variable{
			{VariableIdentifier: "<<user>>", VariableName: "user", VariableValue: "John"},
			{VariableIdentifier: "<<job>>", VariableName: "job", VariableValue: "developer"},
		},
		expectedOutput: "JohndeveloperJohn",
	},
}

func variablesEqualExceptVariableValue(a *[]*Variable, b *[]*Variable) bool {
    if len(*a) != len(*b) {
        return false
    }

    for i := range *a {
        av := reflect.ValueOf((*a)[i]).Elem()
        bv := reflect.ValueOf((*b)[i]).Elem()

        for i := 0; i < av.NumField(); i++ {
            if av.Type().Field(i).Name == "VariableValue" {
                continue
            }
            if av.Field(i).Interface() != bv.Field(i).Interface() {
                return false
            }
        }
    }

    return true
}