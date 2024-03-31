package cmd

import (
	"fmt"

	"github.com/hanskrohn/forge-template/internal/actions"
	"github.com/hanskrohn/forge-template/internal/state"
	"github.com/spf13/cobra"
)

var isFile bool
var isDirectory bool

var fileName string
var templateName string

var Version string = "1.2.0"

var rootCmd = &cobra.Command{
	Use:   "forge-template",
	Short: "create boilerplate code",
	Long:  `create boilerplate code`,
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		s := state.New(state.Unknown)

		actions.DisplayMainMenu(s)


		switch s.Action {
		case state.CreateDirectoryFromTemplate:
			actions.CreateDirectory(nil)
		case state.CreateFileFromTemplate:
			actions.CreateFile(nil)
		case state.CreateDirectoryTemplate:
			actions.CreateTemplate(nil, true)
		case state.CreateFileTemplate:
			actions.CreateTemplate(nil, false)
		case state.DeleteDirectoryTemplate:
			actions.DeleteTemplate(nil, true)
		case state.DeleteFileTemplate:
			actions.DeleteTemplate(nil, false)
		case state.SaveToGithub:
			fallthrough
		default:
			fmt.Println("Unknown action")
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
