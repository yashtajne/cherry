package cmds

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/yashtajne/cherry/utils"
)

func Initalize(project_name string) {
	files, err := os.ReadDir(Project_Work_Directory_Path)
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

	create_build_dir(Project_Build_Debug_Directory_Path)
	create_build_dir(Project_Build_Release_Directory_Path)

	create_lib_dir(Project_Lib_Debug_Directory_Path)
	create_lib_dir(Project_Lib_Release_Directory_Path)

	create_include_dir()
	create_src_dir()

	create_main_file(main_file_name)
	create_gitignore_file()
	create_log_file()
	create_config_file(project_name, compiler)

	fmt.Printf("Project '%s' created successfully!\n", project_name)
}

func create_src_dir() {
	err := os.MkdirAll(Project_Src_Directory_Path, 0755)
	if err != nil {
		fmt.Printf("Error (while creating source directory): %v\n", err)
		return
	}
}

func create_build_dir(build_dir_path string) {
	err := os.MkdirAll(filepath.Join(build_dir_path, "o"), 0755)
	if err != nil {
		fmt.Printf("Error (while creating build directory): %v\n", err)
		return
	}
	err = os.MkdirAll(filepath.Join(build_dir_path, "out"), 0755)
	if err != nil {
		fmt.Printf("Error (while creating source directory): %v\n", err)
		return
	}
}

func create_include_dir() {
	err := os.MkdirAll(Project_Include_Directory_Path, 0755)
	if err != nil {
		fmt.Printf("Error (while creating include directory): %v\n", err)
		return
	}
}

func create_lib_dir(lib_dir_path string) {
	err := os.MkdirAll(lib_dir_path, 0755)
	if err != nil {
		fmt.Printf("Error (while creating lib directory): %v\n", err)
		return
	}
}

func create_log_file() {
	file, err := os.Create(Project_Log_File_Path)
	if err != nil {
		fmt.Printf("Error (while creating log file): %v\n", err)
		return
	}
	defer file.Close()
}

func create_config_file(project_name, compiler string) {
	err := InitConfig(Project_Config_File_Path, project_name, compiler)
	if err != nil {
		fmt.Printf("Error (while creating config file): %v\n", err)
		return
	}
}

func create_main_file(main_file_name string) {
	main_file, err := os.Create(filepath.Join(Project_Src_Directory_Path, main_file_name))
	if err != nil {
		fmt.Printf("Error (while creating .gitignore file): %v\n", err)
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
	gitignore_file, err := os.Create(filepath.Join(Project_Work_Directory_Path, ".gitignore"))
	if err != nil {
		fmt.Printf("Error (while creating .gitignore file): %v\n", err)
		return
	}
	defer gitignore_file.Close()

	gitignore_file.WriteString(DefaultCommandGitignoreFile)
}
