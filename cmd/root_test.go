package cmd

import (
	"testing"
)

func TestRootCommand(t *testing.T) {
    if len(rootCmd.Commands()) != 4 {
        t.Errorf("Expected 1 command, got %d", len(rootCmd.Commands()))
    }

	AssertCommandProperties(t, rootCmd.Commands()[0], "create-directory", []string{"cd", "c-d"}, "Create a directory from a template", "Create a directory from a template")
	AssertCommandProperties(t, rootCmd.Commands()[1], "create-file", []string{"cf", "c-f"}, "Create a file from a template", "Create a file from a template with the given name")
	AssertCommandProperties(t, rootCmd.Commands()[2], "create-template", []string{"ct", "c-t"}, "Create a new template", "Create a new directory or file template")
	AssertCommandProperties(t, rootCmd.Commands()[3], "delete-template", []string{"dt", "d-t"}, "Delete a template", "Delete a directory or file template")
}