package cmd

import (
	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/spf13/cobra"
)


func AddCreateFileCommand(rootCmd *cobra.Command) {
	var createFileCmd = &cobra.Command{
		Use:   "create-file",
		Aliases: []string{"cf", "c-f"},
		Short: "Create a file from a template",
		Long:  `Create a file from a template with the given name`,
		Run: func(cmd *cobra.Command, args []string) {
			actions.CreateFile(newUserInputData(templateName, fileName))
		},
	}
	
	createFileCmd.Flags().StringVarP(&templateName, "templateName", "t", "", "Name of template to use (required)")
	createFileCmd.MarkFlagRequired("templateName")

	createFileCmd.Flags().StringVarP(&fileName, "fileName", "f", "", "Name of the file to create (optional)")

	rootCmd.AddCommand(createFileCmd)
}