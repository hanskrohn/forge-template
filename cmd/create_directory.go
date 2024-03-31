package cmd

import (
	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/spf13/cobra"
)


func AddCreateDirectoryCommand(rootCmd *cobra.Command) {
	var createDirectoryCmd = &cobra.Command{
		Use:   "create-directory",
		Aliases: []string{"cd", "c-d"},
		Short: "Create a directory from a template",
		Long:  `Create a directory from a template`,
		Run: func(cmd *cobra.Command, args []string) {
			actions.CreateDirectory(newUserInputData(templateName, fileName))
		},
	}

	createDirectoryCmd.Flags().StringVarP(&templateName, "templateName", "t", "", "Name of template to use (required)")
	createDirectoryCmd.MarkFlagRequired("templateName")

	rootCmd.AddCommand(createDirectoryCmd)
}