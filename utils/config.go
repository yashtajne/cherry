package utils

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
)

func GetProjectConfig() (*ProjectConfig, error) {
	var config ProjectConfig
	_, err := toml.DecodeFile(Project_Config_File_Path, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func RemovePkgFromConfig(package_name string) {
	data, err := os.ReadFile(Project_Config_File_Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	var project_config ProjectConfig

	err = toml.Unmarshal(data, &project_config)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, pkg := range project_config.Packages {
		if pkg.Name == package_name {
			project_config.Packages = append(project_config.Packages[:i], project_config.Packages[i+1:]...)
			break
		}
	}

	data, err = toml.Marshal(&project_config)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(Project_Config_File_Path, data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func AddPkgToConfig(pkg Pkg) {
	data, err := os.ReadFile(Project_Config_File_Path)
	if err != nil {
		fmt.Println(err)
		return
	}

	var config ProjectConfig

	err = toml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, _package := range config.Packages {
		if _package.Name == pkg.Name {
			config.Packages = append(config.Packages[:i], config.Packages[i+1:]...)
			break
		}
	}

	config.Packages = append(config.Packages, pkg)

	data, err = toml.Marshal(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(Project_Config_File_Path, data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadPackageConfig(packageConfigPath string) (*Pkg, error) {
	data, err := os.ReadFile(packageConfigPath)
	if err != nil {
		return nil, fmt.Errorf("cannot find package config file: %s", err)
	}

	lines := strings.Split(string(data), "\n")
	pkg := &Pkg{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		if !strings.Contains(line, ":") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "Name":
			pkg.Name = value
		case "Description":
			pkg.Description = value
		case "URL":
			pkg.URL = value
		case "Version":
			pkg.Version = value
		case "Libs":
			pkg.Libs = value
		case "Cflags":
			pkg.Cflags = value
		}
	}

	if pkg.Name == "" {
		return nil, errors.New("package name not found in config file")
	}

	return pkg, nil
}

func InitConfig(work_dir_path, project_name, compiler string) error {
	var config ProjectConfig

	config.Project.Name = project_name                       // set project name
	config.Build.IncludeDir = Project_Include_Directory_Path // set include directory path
	config.Build.LibDir = Project_Lib_Debug_Directory_Path   // set lib directory path
	config.Build.Compiler = compiler                         // set compiler path
	config.Build.OS = runtime.GOOS                           // set operating system path
	config.Build.Shell = os.Getenv("SHELL")                  // set shell path
	if config.Build.Shell == "" {
		config.Build.Shell = os.Getenv("ComSpec") // get shell path in windows os
	}

	data, err := toml.Marshal(&config)
	if err != nil {
		return err
	}

	config_file, err := os.OpenFile(Project_Config_File_Path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	_, err = config_file.Write(data)
	if err != nil {
		return err
	}

	defer config_file.Close()

	return nil
}
