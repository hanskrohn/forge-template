package cmd

import (
	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "forge-template",
	Short: "create boilerplate code",
	Long:  `create boilerplate code`,
	Version: "1.0.1", // Make this dynamic
	Run: func(cmd *cobra.Command, args []string) {
		s := state.New()

		actions.DisplayMainMenu(s)

		selectedAction := s.Action

		if(selectedAction == state.CreateProjectTemplate || selectedAction == state.CreateFileTemplate) {
			actions.CreateTemplate(s)
		}else if(selectedAction == state.CreateProjectFromTemplate || selectedAction == state.CreateFileFromTemplate) {
			actions.CreateFileOrDirectory(s)
		}else if (selectedAction == state.DeleteProjectTemplate || selectedAction == state.DeleteFileTemplate) {
			actions.DeleteTemplate(s)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
