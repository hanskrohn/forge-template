package cmd

import (
	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/spf13/cobra"
)

var filename string
var fileTemplateName string

var createFileCmd = &cobra.Command{
	Use:   "create-file",
	Aliases: []string{"cf", "c-f"},
	Short: "Create a file from a template",
	Long:  `Create a file from a template with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		s := state.New(state.CreateFileFromTemplate)
		c := actions.CreateFileOrDirFromTemplateData{
			FileName: 	  filename,
			TemplateName: fileTemplateName,
			Mode:         actions.ModeDefiningVariableNames,
		}
		
		actions.CreateFileOrDirectory(s, &c)
	},
}

func AddCreateFileCommand(rootCmd *cobra.Command) {
	createFileCmd.Flags().StringVarP(&fileTemplateName, "templateName", "t", "", "Name of template to use (required)")
	createFileCmd.MarkFlagRequired("templateName")

	createFileCmd.Flags().StringVarP(&filename, "filename", "f", "", "Name of the file to create (required)")
	createFileCmd.MarkFlagRequired("filename")

	rootCmd.AddCommand(createFileCmd)
}