package cmd

import (
	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/spf13/cobra"
)

var createFileCmd = &cobra.Command{
	Use:   "create-file",
	Aliases: []string{"cf", "c-f"},
	Short: "Create a file from a template",
	Long:  `Create a file from a template with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		actions.CreateFile(newUserInputData(templateName, fileName))
	},
}

func AddCreateFileCommand(rootCmd *cobra.Command) {
	createFileCmd.Flags().StringVarP(&templateName, "templateName", "t", "", "Name of template to use (required)")
	createFileCmd.MarkFlagRequired("templateName")

	createFileCmd.Flags().StringVarP(&fileName, "fileName", "f", "", "Name of the file to create (required)")

	rootCmd.AddCommand(createFileCmd)
}

func newUserInputData(templateName string, fileName string) *actions.UserInputData {
	f := templateName
	if fileName != "" {
		f = fileName
	}

	return &actions.UserInputData {
		TemplateName: templateName,
		FileName: f,
	}
}