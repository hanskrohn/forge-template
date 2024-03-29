package cmd

import (
	"fmt"

	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/spf13/cobra"
)

var filename string
var templateName string

var createFileCmd = &cobra.Command{
	Use:   "create-file",
	Short: "Create a file",
	Long:  `Create a file with the given name`,
	Run: func(cmd *cobra.Command, args []string) {
		s := state.New(state.CreateFileFromTemplate)
		c := actions.CreateFileOrDirFromTemplateData{
			FileName: filename,
			TemplateName: templateName,
			Mode:     actions.ModeDefiningVariableNames,
		}
		fmt.Println(args)
		actions.CreateFileOrDirectory(s, &c)
	},
}

func AddCreateFileCommand(rootCmd *cobra.Command) {
	createFileCmd.Flags().StringVarP(&templateName, "templateName", "t", "", "Name of template to use (required)")
	createFileCmd.MarkFlagRequired("templateName")

	createFileCmd.Flags().StringVarP(&filename, "filename", "f", "", "Name of the file to create (required)")
	createFileCmd.MarkFlagRequired("filename")

	rootCmd.AddCommand(createFileCmd)
}