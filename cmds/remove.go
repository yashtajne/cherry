package cmds

import (
	"fmt"
	"os"
	"path/filepath"

	. "github.com/yashtajne/cherry/utils"
)

func Remove(work_dir_path, package_name string) {
	pkgconfg_path := filepath.Join(work_dir_path, "lib", "pkgconfig")
	pc_files, err := os.ReadDir(pkgconfg_path)
	if err != nil {
		fmt.Printf("Error (while reading pkgconfig directory): %v", err)
		return
	}

	for _, pc_file := range pc_files {
		if !pc_file.IsDir() {
			if pc_file.Name() == (package_name + ".pc") {
				if os.Remove(filepath.Join(pkgconfg_path, pc_file.Name())) != nil {
					fmt.Printf("Error (while removing package): %v", err)
					return
				}
			}
		}
	}

	RemovePkgFromConfig(package_name)
}
