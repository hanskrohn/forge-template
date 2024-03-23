package directories

import (
	"os"

	"github.com/hanskrohn/forge-template/internal/common"
)

func CreateRootConfigDirectory() {
	projectTemplatePath, fileTemplatePath := common.GetImportantDirectories()

	err := os.MkdirAll(projectTemplatePath, 0755)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(fileTemplatePath, 0755)
	if err != nil {
		panic(err)
	}
}