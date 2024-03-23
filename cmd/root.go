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
	Run: func(cmd *cobra.Command, args []string) {
		s := state.New()

		actions.DisplayMainMenu(s)

		selectedAction := s.Action

		if(selectedAction == state.CreateProjectTemplate || selectedAction == state.CreateFileTemplate) {
			actions.CreateTemplate(s)
		}else if(true){

		}

		// switch selectedAction := s.Action; selectedAction {
		// case state.CreateProjectTemplate:
		// 	actions.CreateTemplate(s)
		// case state.CreateFileTemplate:
		// 	actions.CreateTemplate(s)
		// case state.CreateProjectFromTemplate:
		// 	actions.CreateFileOrDirFromTemplate(s)
		// case state.CreateFileFromTemplate:
		// 	actions.CreateFileOrDirFromTemplate(s)
		// case state.DeleteProjectTemplate:
		// 	actions.DeleteFileOrDirTemplate(s)
		// case state.DeleteFileTemplate:
		// 	actions.DeleteFileOrDirTemplate(s)
		// case state.SaveToGithub:
		// 	actions.SaveToGithub(s)
		// }
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
