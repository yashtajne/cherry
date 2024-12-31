package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	. "github.com/yashtajne/cherry/utils"
)

func Run() {
	project_config, err := ReadProjectConfig(ProjectConfigFilePath)
	if err != nil {
		fmt.Printf("Error (reading project config): %v\n", err)
		return
	}

	exe_path := filepath.Join(ProjectWorkDirectoryPath, "build", "out", project_config.Project.Name+".out")

	if _, err := os.Stat(exe_path); err != nil {
		fmt.Printf("Error (executable not found): %v\n", err)
		return
	}

	err = syscall.Exec(exe_path, []string{exe_path}, os.Environ())
	if err != nil {
		fmt.Printf("Error (cannot invoke a system call): %v\n", err)
	}
}
