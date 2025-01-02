package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"github.com/yashtajne/cherry/cmds"
	. "github.com/yashtajne/cherry/utils"
)

const Version string = "1.5.8"

func main() {
	pwd, err := GetWorkDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	Project_Work_Directory_Path = pwd

	Project_Src_Directory_Path = filepath.Join(Project_Work_Directory_Path, "/src")
	Project_Include_Directory_Path = filepath.Join(Project_Work_Directory_Path, "/include")

	Project_Build_Directory_Path = filepath.Join(Project_Work_Directory_Path, "/build")
	Project_Build_Release_Directory_Path = filepath.Join(Project_Build_Directory_Path, "/release")
	Project_Build_Debug_Directory_Path = filepath.Join(Project_Build_Directory_Path, "/debug")

	Project_Lib_Directory_Path = filepath.Join(Project_Work_Directory_Path, "/lib")
	Project_Lib_Release_Directory_Path = filepath.Join(Project_Lib_Directory_Path, "/release")
	Project_Lib_Debug_Directory_Path = filepath.Join(Project_Lib_Directory_Path, "/debug")

	Project_Config_File_Path = filepath.Join(Project_Work_Directory_Path, "/cherry.toml")
	Project_Log_File_Path = filepath.Join(Project_Build_Directory_Path, "/cherry.log")

	app := &cli.App{
		Name:           "cherry",
		Usage:          "A C/C++ Build tool",
		Version:        Version,
		DefaultCommand: "version",
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "Print the version",
				Action: func(c *cli.Context) error {
					cmds.Version(Version)
					return nil
				},
			},
			{
				Name:  "help",
				Usage: "Show help information",
				Action: func(c *cli.Context) error {
					fmt.Println(cli.ShowAppHelp(c))
					return nil
				},
			},
			{
				Name:      "init",
				Usage:     "Initialize a new project",
				ArgsUsage: "<project_name>",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("Error: <project_name> not provided")
					}
					cmds.Initalize(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "make",
				Usage: "Build the project",
				Action: func(c *cli.Context) error {
					cmds.Make()
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add a new package (library) to the project",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("Error: <package_name> not provided")
					}
					cmds.Add(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "Update cherry to the latest version",
				Action: func(c *cli.Context) error {
					cmds.Update()
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "Remove a package (library) from the project",
				Action: func(c *cli.Context) error {
					cmds.Remove(c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "run",
				Usage: "Execute the compiled binary in the current terminal",
				Action: func(c *cli.Context) error {
					cmds.Run()
					return nil
				},
			},
			{
				Name:  "list",
				Usage: "Get a list of installed libraries (packages) in the project",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "vcpkg",
						Usage: "List all vcpkg installed packages",
					},
				},
				Action: func(c *cli.Context) error {
					cmds.List(c.Bool("vcpkg"))
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
