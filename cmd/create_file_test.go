package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddCreateFileCommand(t *testing.T) {
    rootCmd := &cobra.Command{}
    AddCreateFileCommand(rootCmd)

    if len(rootCmd.Commands()) != 1 {
        t.Errorf("Expected 1 command, got %d", len(rootCmd.Commands()))
    }

	AssertCommandProperties(t, rootCmd.Commands()[0], "create-file", []string{"cf", "c-f"}, "Create a file from a template", "Create a file from a template with the given name")
	
	AssertFlagProperties(t, rootCmd.Commands()[0], "templateName", "t", "", "Name of template to use (required)")
	AssertFlagProperties(t, rootCmd.Commands()[0], "fileName", "f", "", "Name of the file to create (optional)")
}