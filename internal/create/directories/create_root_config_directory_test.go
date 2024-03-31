package directories

import (
	"os"
	"testing"

	"github.com/hanskrohn/forge-template/internal/common"
)

func TestCreateRootConfigDirectory(t *testing.T) {
    CreateRootConfigDirectory()

    projectTemplatePath, fileTemplatePath := common.GetImportantDirectories()
    if _, err := os.Stat(projectTemplatePath); os.IsNotExist(err) {
        t.Errorf("Directory %s was not created", projectTemplatePath)
    }
    if _, err := os.Stat(fileTemplatePath); os.IsNotExist(err) {
        t.Errorf("Directory %s was not created", fileTemplatePath)
    }

    os.RemoveAll(projectTemplatePath)
    os.RemoveAll(fileTemplatePath)
}