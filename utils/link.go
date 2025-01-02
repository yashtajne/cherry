package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func Link(project_config *ProjectConfig, o *[]string) error {
	if len(*o) == 0 {
		return fmt.Errorf("No object files")
	}

	// get compiler path
	compiler_path, err := GetCompilerPath(project_config.Build.Compiler)
	if err != nil {
		return err
	}

	// output file path
	output_file_path := filepath.Join(Project_Build_Debug_Directory_Path, "out", project_config.Project.Name)

	// create args for link command
	args := CreateLinkingCommandArgs(project_config)

	*o = append(*o, *args...)               // append args to list of objects
	*o = append(*o, "-o", output_file_path) // append output path to the list of objects

	cmd := exec.Command(project_config.Build.Shell, "-c", fmt.Sprintf("%s %s", compiler_path, strings.Join(*o, " "))) // build commnads

	// fmt.Println(cmd.String()) // print command for debugging

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Linking error: %v\n", err)
		fmt.Printf("Stderr: %s\n", stderr.String())
		return err
	}

	fmt.Printf("\nBuild successfull %s\n", output_file_path)
	return nil
}

func CreateLinkingCommandArgs(project_config *ProjectConfig) *[]string {
	cmd := []string{}

	for _, _package := range project_config.Packages {
		cflags := strings.Replace(strings.ReplaceAll(_package.Libs, "\"", ""), "${libdir}", project_config.Build.LibDir, -1)
		cmd = append(cmd, cflags)
	}

	return &cmd
}
