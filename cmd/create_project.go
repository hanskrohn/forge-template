package cmd

import (
	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/spf13/cobra"
)

var projectTemplateName string

var createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Aliases: []string{"cp", "c-p"},
	Short: "Create a project from a template",
	Long:  `Create a project from a template`,
	Run: func(cmd *cobra.Command, args []string) {
		s := state.New(state.CreateProjectFromTemplate)
		c := actions.CreateFileOrDirFromTemplateData{
			TemplateName: projectTemplateName,
			Mode:         actions.ModeDefiningVariableNames,
		}

		actions.CreateFileOrDirectory(s, &c)
	},
}

func AddCreateProjectCommand(rootCmd *cobra.Command) {
	createProjectCmd.Flags().StringVarP(&projectTemplateName, "templateName", "t", "", "Name of template to use (required)")
	createProjectCmd.MarkFlagRequired("templateName")

	rootCmd.AddCommand(createProjectCmd)
}