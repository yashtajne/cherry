package cmds

import (
	"fmt"
	"os"

	. "github.com/yashtajne/cherry/utils"
)

func Initalize(work_dir_path, project_name string) {
	files, err := os.ReadDir(work_dir_path)
	if err != nil {
		fmt.Printf("Error occured invalid direcotry path: %s\n", err)
		return
	}

	if len(files) != 0 {
		fmt.Println("Error occured directory not empty: cannot initialize this directory")
		return
	}

	err = os.MkdirAll(work_dir_path+"/build/o", 0755)
	if err != nil {
		fmt.Printf("Cannot create build objects directory\nError: %v", err)
		return
	}
	fmt.Println("Created build objects directroy")

	err = os.MkdirAll(work_dir_path+"/build/out", 0755)
	if err != nil {
		fmt.Printf("Cannot create build output directory\nError: %v", err)
		return
	}
	fmt.Println("Created build output directroy")

	err = os.MkdirAll(work_dir_path+"/src", 0755)
	if err != nil {
		fmt.Printf("Cannot create source directory\nError: %v", err)
		return
	}
	fmt.Println("Created source directroy")

	file, err := os.Create("cherry.log")
	if err != nil {
		fmt.Printf("Cannot create cherry.log file\nError: %v", err)
		return
	}
	defer file.Close()
	fmt.Println("Created cherry.log file")

	err = InitConfig(work_dir_path, project_name)
	if err != nil {
		fmt.Printf("Cannot create cherry.toml file\nError: %v", err)
		return
	}
	fmt.Println("Created cherry.toml file")

	fmt.Printf("Successfully initialized %s directory\n", work_dir_path)
}
