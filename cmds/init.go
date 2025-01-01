package cmds

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/yashtajne/cherry/utils"
)

func Initalize(project_name string) {
	files, err := os.ReadDir(ProjectWorkDirectoryPath)
	if err != nil {
		fmt.Printf("Error occured invalid direcotry path: %s\n", err)
		return
	}

	if len(files) != 0 {
		fmt.Println("Error occured directory not empty: cannot initialize this directory")
		return
	}

	main_file_name := ""
	compiler := ""

	fmt.Println("What type of project would you like to create?")
	fmt.Println("(A) [C project]")
	fmt.Println("(B) [C++ project]")
	fmt.Println("(Q) Quit")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter your choice: ")

	char, err := reader.ReadByte()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch char {
	case 'A', 'a':
		main_file_name = "main.c"
		compiler = "gcc"
	case 'B', 'b':
		main_file_name = "main.cpp"
		compiler = "g++"
	case 'Q', 'q':
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid choice. Exiting...")
		return
	}

	create_config_file(project_name, compiler)
	create_log_file()
	create_build_dir()
	create_include_dir()
	create_lib_dir()
	create_src_dir()
	create_main_file(main_file_name)
	create_gitignore_file()
	fmt.Printf("Project '%s' initialized successfully!\n", project_name)
}

func create_src_dir() {
	err := os.MkdirAll(ProjectSrcDirectoryPath, 0755)
	if err != nil {
		fmt.Printf("Error (while creating source directory): %v\n", err)
		return
	}
}

func create_build_dir() {
	err := os.MkdirAll(filepath.Join(ProjectBuildDirectoryPath, "o"), 0755)
	if err != nil {
		fmt.Printf("Error (while creating build directory): %v\n", err)
		return
	}
	err = os.MkdirAll(filepath.Join(ProjectBuildDirectoryPath, "out"), 0755)
	if err != nil {
		fmt.Printf("Error (while creating source directory): %v\n", err)
		return
	}
}

func create_include_dir() {
	err := os.MkdirAll(ProjectIncludeDirectoryPath, 0755)
	if err != nil {
		fmt.Printf("Error (while creating include directory): %v\n", err)
		return
	}
}

func create_lib_dir() {
	err := os.MkdirAll(filepath.Join(ProjectLibDirectoryPath, "pkgconfig"), 0755)
	if err != nil {
		fmt.Printf("Error (while creating lib directory): %v\n", err)
		return
	}
}

func create_log_file() {
	file, err := os.Create(ProjectLogFilePath)
	if err != nil {
		fmt.Printf("Error (while creating log file): %v\n", err)
		return
	}
	defer file.Close()
}

func create_config_file(project_name, compiler string) {
	err := InitConfig(ProjectWorkDirectoryPath, project_name, compiler)
	if err != nil {
		fmt.Printf("Error (while creating config file): %v\n", err)
		return
	}
}

func create_main_file(main_file_name string) {
	main_file_path := filepath.Join(ProjectSrcDirectoryPath, main_file_name)

	main_file, err := os.Create(main_file_path)
	if err != nil {
		fmt.Printf("Error (while creating %s file): %v\n", main_file_path, err)
		return
	}

	defer main_file.Close()

	if main_file_name == "main.c" {
		main_file.WriteString(DefaultMainCFile)
	} else if main_file_name == "main.cpp" {
		main_file.WriteString(DefaultMainCPPFile)
	}
}

func create_gitignore_file() {
	gitignore_file, err := os.Create(filepath.Join(ProjectWorkDirectoryPath, ".gitignore"))
	if err != nil {
		fmt.Printf("Error (while creating .gitignore file): %v\n", err)
		return
	}

	defer gitignore_file.Close()

	gitignore_file.WriteString(DefaultCommandGitignoreFile)
}
