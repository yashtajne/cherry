package cmds

import (
	"fmt"

	. "github.com/yashtajne/cherry/utils"
)

// Lists the packages installed in the current project directory
func List() {
	project_config, err := GetProjectConfig()
	if err != nil {
		fmt.Printf("Error (reading project config): %v", err)
		return
	}

	for i, _package := range project_config.Packages {
		fmt.Printf("%d. %s v%s\n", i+1, _package.Name, _package.Version)
	}
}
