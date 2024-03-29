package cmd

import (
	"fmt"

	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/spf13/cobra"
)

var isFile bool
var isProject bool

var templateName string

var createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Aliases: []string{"ct", "c-t"},
	Short: "Create a new template",
	Long:  `Create a new project or file template`,
	Run: func(cmd *cobra.Command, args []string) {
		if isFile == isProject {
            fmt.Println("You must specify either --file or --project.")
            return
        }

		var s *state.State
		if isFile {
			s = state.New(state.CreateFileTemplate)
		}else{
			s = state.New(state.CreateProjectTemplate)
		}

		c := actions.CreateTemplateData{
			TemplateName: templateName,
			Mode:         actions.ModeDefiningContent,
		}
		
		actions.CreateTemplate(s, &c)
	},
}

func AddCreateTemplateCommand(rootCmd *cobra.Command) {
	createTemplateCmd.Flags().BoolVarP(&isFile, "file", "f", false, "Create a file template")
    createTemplateCmd.Flags().BoolVarP(&isProject, "project", "p", false, "Create a project template")

	createTemplateCmd.Flags().StringVarP(&templateName, "templateName", "t", "", "Template name (required)")
	createTemplateCmd.MarkFlagRequired("templateName")

	rootCmd.AddCommand(createTemplateCmd)
}