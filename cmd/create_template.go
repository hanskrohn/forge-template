package cmd

import (
	"fmt"

	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/spf13/cobra"
)


func AddCreateTemplateCommand(rootCmd *cobra.Command) {
	var createTemplateCmd = &cobra.Command{
		Use:   "create-template",
		Aliases: []string{"ct", "c-t"},
		Short: "Create a new template",
		Long:  `Create a new directory or file template`,
		Run: func(cmd *cobra.Command, args []string) {
			if isFile == isDirectory {
				fmt.Println("You must specify either --file or --directory.")
				return
			}
	
			if isFile {
				actions.CreateTemplate(newUserInputData(templateName, fileName), false)
			}else{
				actions.CreateTemplate(newUserInputData(templateName, fileName), true)
			}
		},
	}
	
	createTemplateCmd.Flags().BoolVarP(&isFile, "file", "f", false, "Create a file template")
    createTemplateCmd.Flags().BoolVarP(&isDirectory, "directory", "d", false, "Create a directory template")

	createTemplateCmd.Flags().StringVarP(&templateName, "templateName", "t", "", "Template name (required)")
	createTemplateCmd.MarkFlagRequired("templateName")

	rootCmd.AddCommand(createTemplateCmd)
}