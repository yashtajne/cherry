package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/yashtajne/cherry/cmds"
	. "github.com/yashtajne/cherry/utils"
)

const Version string = "1.1.1"

func main() {
	pwd, err := GetWorkDir()
	if err != nil {
		log.Fatal(err)
	}

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
					cmds.Initalize(pwd, c.Args().Get(0))
					return nil
				},
			},
			{
				Name:  "bake",
				Usage: "Build the project",
				Action: func(c *cli.Context) error {
					cmds.Make(pwd)
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
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
