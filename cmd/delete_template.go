package cmd

import (
	"fmt"

	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/spf13/cobra"
)


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

		var s *state.State
		if isFile {
			s = state.New(state.DeleteFileTemplate)
		}else{
			s = state.New(state.DeleteDirectoryTemplate)
		}
		
		actions.DeleteTemplate(s, templateName)
	},
}

func AddDeleteTemplateCommand(rootCmd *cobra.Command) {
	deleteTemplateCmd.Flags().BoolVarP(&isFile, "file", "f", false, "Create a file template")
    deleteTemplateCmd.Flags().BoolVarP(&isDirectory, "directory", "p", false, "Create a directory template")

	deleteTemplateCmd.Flags().StringVarP(&templateName, "templateName", "t", "", "Template name (required)")
	deleteTemplateCmd.MarkFlagRequired("templateName")

	rootCmd.AddCommand(deleteTemplateCmd)
}