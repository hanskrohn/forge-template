package cmd

import (
	"fmt"

	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/spf13/cobra"
)



func AddDeleteTemplateCommand(rootCmd *cobra.Command) {
	var deleteTemplateCmd = &cobra.Command{
		Use:   "delete-template",
		Aliases: []string{"dt", "d-t"},
		Short: "Delete a template",
		Long:  `Delete a directory or file template`,
		Run: func(cmd *cobra.Command, args []string) {
			if isFile == isDirectory {
				fmt.Println("You must specify either --file or --Directory.")
				return
			}
	
			if isFile {
				actions.DeleteTemplate(newUserInputData(templateName, fileName), false)
			}else{
				actions.DeleteTemplate(newUserInputData(templateName, fileName), true)
			}
		},
	}
	
	deleteTemplateCmd.Flags().BoolVarP(&isFile, "file", "f", false, "Create a file template")
    deleteTemplateCmd.Flags().BoolVarP(&isDirectory, "directory", "d", false, "Create a directory template")

	deleteTemplateCmd.Flags().StringVarP(&templateName, "templateName", "t", "", "Template name (required)")
	deleteTemplateCmd.MarkFlagRequired("templateName")

	rootCmd.AddCommand(deleteTemplateCmd)
}