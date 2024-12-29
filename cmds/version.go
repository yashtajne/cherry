package cmds

import (
	"fmt"

	. "github.com/yashtajne/cherry/utils"
)

func Version(version string) {

	// gcc version
	gcc_version, err := GetCompilerVersion("gcc")
	if err != nil {
		fmt.Printf("Error (getting gcc version): %v", err)
		return
	}

	// g++ version
	gpp_version, err := GetCompilerVersion("g++")
	if err != nil {
		fmt.Printf("Error (getting g++ version): %v", err)
		return
	}

	fmt.Println("version: ", version)
	fmt.Println(gcc_version)
	fmt.Println(gpp_version)
}
