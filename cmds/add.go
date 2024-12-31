package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/yashtajne/cherry/utils"
)

func Add(work_dir_path, package_name string) {

	// get $VCPKG_ROOT path
	vcpkg_root := os.Getenv("VCPKG_ROOT")
	if vcpkg_root == "" {
		fmt.Println("Error: $VCPKG_ROOT is not set, make sure vcpkg is installed and $VCPKG_ROOT PATH is set")
		return
	}

	vcpkg_packages_dir := filepath.Join(vcpkg_root, "packages")

	packages, err := os.ReadDir(vcpkg_packages_dir)
	if err != nil {
		fmt.Printf("Error (reading vcpkg packages): %v", err)
		return
	}

	for _, _package := range packages {
		_package_name := _package.Name()
		if package_name == strings.Split(_package_name, "_")[0] {
			if err := _add(work_dir_path, filepath.Join(vcpkg_packages_dir, _package_name)); err != nil {
				fmt.Printf("Error (while adding vcpkg package): %v", err)
				return
			}
		}
	}

	pkg, err := ReadPackageConfig(filepath.Join(work_dir_path, "lib", "pkgconfig", package_name+".pc"))
	if err != nil {
		fmt.Printf("Error (while reading package config): %v", err)
		return
	}

	AddPkgToConfig(filepath.Join(work_dir_path, "cherry.toml"), *pkg)
}

func _add(work_dir_path, package_dir_path string) error {
	package_include_dir_path := filepath.Join(package_dir_path, "include") // package include directory
	package_lib_dir_path := filepath.Join(package_dir_path, "lib")         // package lib directory

	work_include_dir_path := filepath.Join(work_dir_path, "include") // project include directory
	work_lib_dir_path := filepath.Join(work_dir_path, "lib")         // project lib directory

	// Check if the include directory exists, and create it if not
	if _, err := os.Stat(work_include_dir_path); os.IsNotExist(err) {
		if err := os.MkdirAll(work_include_dir_path, os.ModePerm); err != nil {
			return fmt.Errorf("error creating include directory: %w", err)
		}
	}

	// Check if the lib directory exists, and create it if not
	if _, err := os.Stat(work_lib_dir_path); os.IsNotExist(err) {
		if err := os.MkdirAll(work_lib_dir_path, os.ModePerm); err != nil {
			return fmt.Errorf("error creating lib directory: %w", err)
		}
	}

	include_contents, err := os.ReadDir(package_include_dir_path)
	if err != nil {
		return err
	}

	lib_contents, err := os.ReadDir(package_lib_dir_path)
	if err != nil {
		return err
	}

	for _, include_content := range include_contents {
		include_content_name := include_content.Name()
		if include_content.IsDir() {
			if err := CopyDirectory(
				filepath.Join(package_include_dir_path, include_content_name),
				filepath.Join(work_include_dir_path, include_content_name),
			); err != nil {
				return err
			}
		}
	}

	for _, lib_content := range lib_contents {
		lib_content_name := lib_content.Name()

		if lib_content.IsDir() {
			if err := CopyDirectory(
				filepath.Join(package_lib_dir_path, lib_content_name),
				filepath.Join(work_lib_dir_path, lib_content_name),
			); err != nil {
				return fmt.Errorf("error copying directory %s: %w", lib_content_name, err)
			}
		} else {
			if err := CopyFile(
				filepath.Join(package_lib_dir_path, lib_content_name),
				filepath.Join(work_lib_dir_path, lib_content_name),
			); err != nil {
				return fmt.Errorf("error copying file %s: %w", lib_content_name, err)
			}
		}
	}

	return nil
}
