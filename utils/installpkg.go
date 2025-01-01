package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func PromptWouldYouLikeToInstallPackage(package_name string) bool {
	fmt.Println("Package not Installed locally.")
	fmt.Println("Would you like me to install this package for you?")
	fmt.Println("(Y) YES (N) NO")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your choice: ")

	char, err := reader.ReadByte()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return false
	}

	switch char {
	case 'Y', 'y':
		return true
	case 'N', 'n':
		return false
	default:
		fmt.Println("Invalid option. Exiting...")
		return false
	}
}

func InstallPkg(vcpkg_root, package_name string) {
	cmd := exec.Command(filepath.Join(vcpkg_root, "vcpkg"), "install", package_name)
	fmt.Println(cmd.String())
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: (while installing %s) %v\n", package_name, err)
		fmt.Printf("--> %s\n", stderr.String())
		return
	}
	fmt.Println("Succesfully installed ", package_name)
}
