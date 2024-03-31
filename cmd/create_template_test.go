package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestAddCreateTemplateCommand(t *testing.T) {
    rootCmd := &cobra.Command{}
    AddCreateTemplateCommand(rootCmd)

    if len(rootCmd.Commands()) != 1 {
        t.Errorf("Expected 1 command, got %d", len(rootCmd.Commands()))
    }

	AssertCommandProperties(t, rootCmd.Commands()[0], "create-template", []string{"ct", "c-t"}, "Create a new template", "Create a new directory or file template")
	
	AssertFlagProperties(t, rootCmd.Commands()[0], "file", "f", "false", "Create a file template")
	AssertFlagProperties(t, rootCmd.Commands()[0], "directory", "d", "false", "Create a directory template")
	AssertFlagProperties(t, rootCmd.Commands()[0], "templateName", "t", "", "Template name (required)")
}