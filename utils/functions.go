package utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var Re = regexp.MustCompile(`_.*`)

func GetWorkDir() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Error (cannot get work directory): %v", err)
	}
	return pwd, nil
}

func GetCompilerPath(compiler string) (string, error) {
	compiler_path, err := exec.LookPath(compiler)
	if err != nil {
		return "", err
	}
	return compiler_path, nil
}

func GetCompilerVersion(compiler string) (string, error) {
	compiler_path, err := GetCompilerPath(compiler)
	if err != nil {
		fmt.Printf("Error (getting gcc path): %v", err)
		return "", err
	}
	compiler_version, err := exec.Command(compiler_path, "--version").Output()
	if err != nil {
		return "", err
	}
	version := strings.SplitN(string(compiler_version), "\n", 2)[0]
	return version, nil
}

func IsCompiled(o_files []os.DirEntry, src_file_name string) bool {
	for _, o_file := range o_files {
		if o_file.Name() == (strings.Replace(src_file_name, ".", "_", -1) + ".o") {
			return true
		}
	}
	return false
}

func SrcFileExist(src_files []os.DirEntry, o_file_name string) bool {
	for _, src_file := range src_files {
		src_file_ext := filepath.Ext(src_file.Name())
		if filepath.Base(o_file_name) == filepath.Base(src_file.Name())+strings.Replace(src_file_ext, ".", "_", -1) {
			return true
		}
	}
	return false
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	info, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, info.Mode())
}

func CopyDirectory(srcDir, dstDir string) error {
	return filepath.Walk(srcDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dstDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		} else {
			return CopyFile(path, targetPath)
		}
	})
}
