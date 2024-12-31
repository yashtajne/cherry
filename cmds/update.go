package cmds

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Update() {
	// Path to the `go` executable
	goPath, err := exec.LookPath("go")
	if err != nil {
		fmt.Printf("Error (finding go binary): %v\n", err)
		return
	}

	// Command arguments
	args := []string{goPath, "install", "github.com/yashtajne/cherry@latest"}

	// Execute the command
	err = syscall.Exec(goPath, args, os.Environ())
	if err != nil {
		fmt.Printf("Error (cannot invoke syscall): %v\n", err)
	}
}
