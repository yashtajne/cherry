package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	. "github.com/yashtajne/cherry/utils"
)

func Run(work_dir_path string) {
	project_config, err := ReadConfig(work_dir_path + "/cherry.toml")
	if err != nil {
		fmt.Printf("Error (reading project config): %v\n", err)
		return
	}

	exe_path := filepath.Join(work_dir_path, "build", "out", project_config.Project.Name+".out")

	if _, err := os.Stat(exe_path); err != nil {
		fmt.Printf("Error (executable not found): %v\n", err)
		return
	}

	err = syscall.Exec(exe_path, []string{exe_path}, os.Environ())
	if err != nil {
		fmt.Printf("Error (cannot invoke a system call): %v\n", err)
	}
}