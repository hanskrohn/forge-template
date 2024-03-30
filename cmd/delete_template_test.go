package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddDeleteTemplateCommand(t *testing.T) {
    rootCmd := &cobra.Command{}
    AddDeleteTemplateCommand(rootCmd)

    if len(rootCmd.Commands()) != 1 {
        t.Errorf("Expected 1 command, got %d", len(rootCmd.Commands()))
    }

	AssertCommandProperties(t, rootCmd.Commands()[0], "delete-template", []string{"dt", "d-t"}, "Delete a template", "Delete a directory or file template")
	
	AssertFlagProperties(t, rootCmd.Commands()[0], "file", "f", "false", "Create a file template")
	AssertFlagProperties(t, rootCmd.Commands()[0], "directory", "d", "false", "Create a directory template")
	AssertFlagProperties(t, rootCmd.Commands()[0], "templateName", "t", "", "Template name (required)")
}