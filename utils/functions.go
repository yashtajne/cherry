package utils

import (
	"fmt"
	"io"
	"log"
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

func CopyDir(src, dst string) error {
	var err error
	var fds []os.DirEntry

	fds, err = os.ReadDir(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	for _, fd := range fds {
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())

		if fd.IsDir() {
			err = CopyDir(srcfp, dstfp)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcfp, dstfp)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	if _, err := os.Stat(dst); err == nil {
		return nil
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return nil
}

func compareFiles(src, dst string) bool {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Open(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	srcContents, err := io.ReadAll(srcFile)
	if err != nil {
		log.Fatal(err)
	}

	dstContents, err := io.ReadAll(dstFile)
	if err != nil {
		log.Fatal(err)
	}

	return string(dstContents) == string(srcContents)
}

func RemoveFiles(srcDir, dstDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relPath, err := filepath.Rel(srcDir, path)
			if err != nil {
				return err
			}

			dstPath := filepath.Join(dstDir, relPath)

			if compareFiles(path, dstPath) {
				err = os.Remove(dstPath)
				if err != nil {
					return err
				}

				fmt.Printf("Removed '%s'\n", dstPath)
			}
		}

		return nil
	})
}
