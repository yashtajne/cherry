package cmds

import (
	"fmt"
	"os"

	. "github.com/yashtajne/cherry/utils"
)

func Add(work_dir_path, package_name string) {
	if _, err := os.Stat(work_dir_path); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Directory '%s' does not exist\n", work_dir_path)
		} else {
			fmt.Printf("Error checking directory: %v\n", err)
		}
	}

	lib_dir_path := work_dir_path + "/lib"
	include_dir_path := work_dir_path + "/include"

	vcpkg_root, exists := os.LookupEnv("VCPKG_ROOT")
	if !exists {
		fmt.Println("Error $VCPKG_ROOT is not set")
		return
	}

	vcpkg_packages_dir := vcpkg_root + "/packages"
	vcpkg_packages, err := os.ReadDir(vcpkg_packages_dir)
	if err != nil {
		fmt.Printf("Error finding packages: %v\n", err)
		return
	}

	for _, vcpkg_package := range vcpkg_packages {
		if Re.ReplaceAllString(vcpkg_package.Name(), "") == package_name && vcpkg_package.IsDir() {
			vcpkg_package_include_dir := vcpkg_packages_dir + "/" + vcpkg_package.Name() + "/include"
			vcpkg_package_lib_dir := vcpkg_packages_dir + "/" + vcpkg_package.Name() + "/lib"

			if err := CopyDir(vcpkg_package_include_dir, include_dir_path); err != nil {
				fmt.Printf("Error while adding this package: %v\n", err)
				return
			}
			if err := CopyDir(vcpkg_package_lib_dir, lib_dir_path); err != nil {
				fmt.Printf("Error while adding this package: %v\n", err)
				return
			}

			vcpkg_package_pc_file_path := work_dir_path + "/lib/pkgconfig/" + package_name + ".pc"
			pkg, err := ReadPackageConfig(vcpkg_package_pc_file_path)
			if err != nil {
				fmt.Printf("Error while reading this package config: %v\n", err)
				return
			}

			AddPkgToConfig(work_dir_path+"/cherry.toml", *pkg)

			fmt.Printf("Succesfully added package: %s\n", package_name)
			return
		}
	}

	fmt.Printf("Package not installed locally: %s\nInstall the package using this command: vcpkg install %s\n", package_name, package_name)
}