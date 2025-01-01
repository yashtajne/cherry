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

const Version string = "1.3.2"

func main() {
	pwd, err := GetWorkDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	ProjectWorkDirectoryPath = pwd
	ProjectSrcDirectoryPath = filepath.Join(ProjectWorkDirectoryPath, "/src")
	ProjectBuildDirectoryPath = filepath.Join(ProjectWorkDirectoryPath, "/build")
	ProjectIncludeDirectoryPath = filepath.Join(ProjectWorkDirectoryPath, "/include")
	ProjectLibDirectoryPath = filepath.Join(ProjectWorkDirectoryPath, "/lib")
	ProjectConfigFilePath = filepath.Join(ProjectWorkDirectoryPath, "/cherry.toml")
	ProjectLogFilePath = filepath.Join(ProjectWorkDirectoryPath, "/cherry.log")

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
					cmds.Help("")
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
					cmds.Add(pwd, c.Args().Get(0))
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
					cmds.Remove(pwd, c.Args().Get(0))
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
