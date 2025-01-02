package cmds

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/yashtajne/cherry/utils"
)

func Add(package_name string) {

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

	package_path := ""
	for _, _package := range packages {
		_package_name := _package.Name()
		if package_name == strings.Split(_package_name, "_")[0] {
			package_path = filepath.Join(vcpkg_packages_dir, _package_name)
			break
		}
	}

	if err := _add(package_path); err != nil {
		if err.Error() == PACKAGE_NOT_INSTALLED_LOCALLY {
			if PromptWouldYouLikeToInstallPackage(package_name) {
				InstallPkg(vcpkg_root, package_name)
				return
			}
			fmt.Printf("Okay, Exiting...")
			return
		}
		fmt.Printf("Error (while adding vcpkg package): %v", err)
		return
	}

	pkg, err := ReadPackageConfig(filepath.Join(Project_Lib_Directory_Path, "pkgconfig", package_name+".pc"))
	if err != nil {
		fmt.Printf("Error (while reading package config): %v", err)
		return
	}

	AddPkgToConfig(*pkg)
}

func _add(package_dir_path string) error {
	// Check if the package directory exists, if not return error
	if !DirExists(package_dir_path) {
		return fmt.Errorf(PACKAGE_NOT_INSTALLED_LOCALLY)
	}

	package_include_dir_path := filepath.Join(package_dir_path, "include")        // package include directory
	package_release_lib_dir_path := filepath.Join(package_dir_path, "lib")        // package release lib directory
	package_debug_lib_dir_path := filepath.Join(package_dir_path, "debug", "lib") // package release lib directory

	// get include files
	include_contents, err := os.ReadDir(package_include_dir_path)
	if err != nil {
		return err
	}

	// get release libraries files
	release_lib_contents, err := os.ReadDir(package_release_lib_dir_path)
	if err != nil {
		return err
	}

	// get debug libraries files
	debug_lib_contents, err := os.ReadDir(package_debug_lib_dir_path)
	if err != nil {
		return err
	}

	// copy include files
	for _, include_content := range include_contents {
		include_content_name := include_content.Name()
		if include_content.IsDir() {
			if err := CopyDirectory(
				filepath.Join(package_include_dir_path, include_content_name),
				filepath.Join(Project_Include_Directory_Path, include_content_name),
			); err != nil {
				return err
			}
		}
	}

	// copy libraries function
	copy_libraries := func(libraries []os.DirEntry, src_path, dest_path string) error {
		for _, lib := range libraries {
			lib_content_name := lib.Name()

			if lib.IsDir() && lib.Name() == "pkgconfig" {
				if err := CopyDirectory(
					filepath.Join(src_path, lib_content_name),
					filepath.Join(Project_Lib_Directory_Path, lib_content_name),
				); err != nil {
					return fmt.Errorf("error copying directory %s: %w", lib_content_name, err)
				}
			} else {
				if err := CopyFile(
					filepath.Join(src_path, lib_content_name),
					filepath.Join(dest_path, lib_content_name),
				); err != nil {
					return fmt.Errorf("error copying file %s: %w", lib_content_name, err)
				}
			}
		}
		return nil
	}

	err = copy_libraries(release_lib_contents, package_release_lib_dir_path, Project_Lib_Release_Directory_Path)
	err = copy_libraries(debug_lib_contents, package_debug_lib_dir_path, Project_Lib_Debug_Directory_Path)
	if err != nil {
		return err
	}

	return nil
}
