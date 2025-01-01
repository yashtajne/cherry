package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func Compile(project_config *ProjectConfig, src_file_path, output_file_path string) error {

	// get the compiler path based on the source file extention
	var compiler string
	if filepath.Ext(src_file_path) == ".c" {
		compiler = "gcc" // gcc for compiling c files
	} else if filepath.Ext(src_file_path) == ".cpp" {
		compiler = "g++" // g++ for compiling c++ files
	}

	// get compiler path
	var compiler_path, err = exec.LookPath(compiler)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	// create args for compile command
	args := CreateCompileCommandArgs(project_config)

	// append source file and output path
	*args = append(*args, "-Wall", "-c", src_file_path, "-o", output_file_path)
	cmd := exec.Command(compiler_path, *args...)

	// print command for debugging
	// fmt.Println(cmd.String())

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Compilation error: %v\n", err)
		fmt.Printf("Stderr: %s\n", stderr.String())
		return err
	}

	fmt.Printf("Compiled %s\n", filepath.Base(src_file_path))
	return nil
}

func CreateCompileCommandArgs(project_config *ProjectConfig) *[]string {
	cmd := []string{}

	for _, _package := range project_config.Packages {
		cflags := strings.Replace(strings.ReplaceAll(_package.Cflags, "\"", ""), "${includedir}", project_config.Build.IncludeDir, -1)
		cmd = append(cmd, cflags)
	}

	return &cmd
}
