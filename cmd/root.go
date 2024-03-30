package cmd

import (
	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/spf13/cobra"
)

var Version string = "1.1.2"

var rootCmd = &cobra.Command{
	Use:   "forge-template",
	Short: "create boilerplate code",
	Long:  `create boilerplate code`,
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		s := state.New(state.Unknown)

		actions.DisplayMainMenu(s)

		selectedAction := s.Action
		if(selectedAction == state.CreateFileFromTemplate) {
			actions.CreateFile(nil)
			return
		}
		if(selectedAction == state.CreateDirectoryFromTemplate) {
			actions.CreateDirectory(nil)
			return
		}

		if(selectedAction == state.CreateDirectoryTemplate || selectedAction == state.CreateFileTemplate) {
			actions.CreateTemplate(s, nil)
		}else if(selectedAction == state.CreateDirectoryFromTemplate || selectedAction == state.CreateFileFromTemplate) {
			// actions.CreateFileOrDirectory(s, nil)
		}else if (selectedAction == state.DeleteDirectoryTemplate || selectedAction == state.DeleteFileTemplate) {
			actions.DeleteTemplate(s, "")
		}
	},
}

func init() {
	AddCreateFileCommand(rootCmd)
	AddCreateDirectoryCommand(rootCmd)
	AddCreateTemplateCommand(rootCmd)
	AddDeleteTemplateCommand(rootCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
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
