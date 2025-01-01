package cmds

import (
	"bytes"
	"fmt"
	"os/exec"

	. "github.com/yashtajne/cherry/utils"
)

// Lists the packages installed in the current project directory
func List(vcpkg bool) {
	if vcpkg {
		cmd := exec.Command("vcpkg", "list")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error (while getting installed packages list): %v\n", err)
			fmt.Printf("----> %s\n", stderr.String())
			return
		}
		fmt.Println(cmd.Stdout)
	} else {
		project_config, err := GetProjectConfig()
		if err != nil {
			fmt.Printf("Error (reading project config): %v", err)
			return
		}

		for i, _package := range project_config.Packages {
			fmt.Printf("%d. %s v%s\n", i+1, _package.Name, _package.Version)
		}
	}
}
