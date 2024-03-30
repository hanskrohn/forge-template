package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddCreateDirectoryCommand(t *testing.T) {
    rootCmd := &cobra.Command{}
    AddCreateDirectoryCommand(rootCmd)

    if len(rootCmd.Commands()) != 1 {
        t.Errorf("Expected 1 command, got %d", len(rootCmd.Commands()))
    }

	AssertCommandProperties(t, rootCmd.Commands()[0], "create-directory", []string{"cd", "c-d"}, "Create a directory from a template", "Create a directory from a template")
	
    AssertFlagProperties(t, rootCmd.Commands()[0], "templateName", "t", "", "Name of template to use (required)")
}