package main

import (
	"fmt"
	"os"

	"github.com/yashtajne/cherry/cmds"
)

const Version string = "1.0.1"

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error (cannot get work directory): %v", err)
	}

	if len(os.Args) < 2 {
		fmt.Printf("Error: <command> not provided")
		return
	}

	switch os.Args[1] {
	case "version":
		cmds.Version(Version)
		break
	case "help":
		cmds.Help("")
		break
	case "init":
		if len(os.Args) < 3 {
			fmt.Printf("Error: <project_name> not provided")
			return
		}
		cmds.Initalize(pwd, os.Args[2])
		break
	case "make":
		cmds.Make(pwd)
		break
	case "add":
		if len(os.Args) < 3 {
			fmt.Printf("Error: <package_name> not provided")
			return
		}
		cmds.Add(pwd, os.Args[2])
		break
	case "remove":
		if len(os.Args) < 3 {
			fmt.Printf("Error: <package_name> not provided")
			return
		}
		cmds.Remove(pwd, os.Args[2])
		break
	case "run":
		cmds.Run(pwd)
		break
	default:
		fmt.Print("Invalid command.")
		break
	}
}
