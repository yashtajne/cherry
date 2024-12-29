package cmds

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	. "github.com/yashtajne/cherry/utils"
)

func Make(work_dir_path string) {
	// read project config file
	project_config, err := ReadConfig(work_dir_path + "/cherry.toml")
	if err != nil {
		fmt.Printf("Error (reading project config): %v", err)
		return
	}

	// file paths
	build_dir_o_path := work_dir_path + "/build/o"
	build_dir_out_path := work_dir_path + "/build/out"
	src_dir_path := work_dir_path + "/src"

	// open log file
	log_file, err := os.OpenFile(work_dir_path+"/cherry.log", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("Error (reading log file): %v", err)
		return
	}

	// check if log file is empty
	if log_file_info, err := log_file.Stat(); err != nil || log_file_info.Size() == 0 {
		if err != nil {
			fmt.Printf("Error (reading log file stats): %v\n", err)
			return
		} else {
			fmt.Println("Empty log file")
		}
	}

	// store contents of the log fle in a map
	old_log := make(map[string]string)
	scanner := bufio.NewScanner(log_file)
	for scanner.Scan() {
		log_line := scanner.Text()
		end_index := strings.Index(log_line, "]")
		date_str := log_line[1:end_index]
		file_name := strings.TrimSpace(log_line[end_index+1:])
		old_log[file_name] = date_str
	}

	// handle scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error (scanner): %v", err)
	}

	// get a list of object files
	o_files, err := os.ReadDir(build_dir_o_path)
	if err != nil {
		fmt.Printf("Error (reading object files): %v", err)
		return
	}

	// get a list of source files
	src_files, err := os.ReadDir(src_dir_path)
	if err != nil {
		fmt.Printf("Error (reading source files): %v", err)
		return
	}

	new_log_file := bytes.NewBufferString("") // create new log file buffer

	for _, src_file := range src_files {
		if !src_file.IsDir() {
			src_file_name := src_file.Name()                      // source file name
			src_file_ext := filepath.Ext(src_file_name)           // source file extentiom
			if src_file_ext != C_EXT && src_file_ext != CPP_EXT { // if file is not c or cpp then continue to next itteration
				continue
			}
			src_file_info, err := src_file.Info() // file stats
			if err != nil {
				fmt.Printf("Error (reading src file stats): %v", err)
				return
			}

			src_file_mod_time_before := old_log[src_file_name]                   // search for mod time entry in the log file
			src_file_mod_time_now := src_file_info.ModTime().Format(TIME_FORMAT) // currently read mod time entry

			if !IsCompiled(o_files, src_file_name) { // file not compiled
				if !_compile(project_config, src_dir_path, src_file_name, src_file_ext, build_dir_o_path) {
					return
				}
				new_log_file.WriteString(fmt.Sprintf("[%s] %s\n", src_file_mod_time_now, src_file_name))
			} else if src_file_mod_time_before != src_file_mod_time_now { // if file is modified
				if !_compile(project_config, src_dir_path, src_file_name, src_file_ext, build_dir_o_path) {
					return
				}
				new_log_file.WriteString(fmt.Sprintf("[%s] %s\n", src_file_mod_time_now, src_file_name))
			} else { // no changes to the file
				new_log_file.WriteString(fmt.Sprintf("[%s] %s\n", src_file_mod_time_before, src_file_name))
			}
		}
	}

	// check if all compiled files have their src files in the source folder if not then delete object file
	for _, o_file := range o_files {
		if !SrcFileExist(src_files, o_file.Name()) {
			os.Remove(o_file.Name())
		}
	}

	// delete all contents of the log file
	err = log_file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// reset the line pointer to the start of the file
	_, err = log_file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// update log file
	log_file.WriteString(new_log_file.String())

	if !_link(project_config, build_dir_o_path, build_dir_out_path) {
		return
	}
}

func _compile(project_config *ProjectConfig, src_dir_path, src_file_name, src_file_ext, build_dir_o_path string) bool {
	if err := Compile(
		project_config, // project config
		filepath.Join(src_dir_path, src_file_name), // source file
		filepath.Join(build_dir_o_path, strings.Replace(src_file_name, src_file_ext, "", -1)+strings.Replace(src_file_ext, ".", "_", -1)+O_EXT), // output file path
	); err != nil {
		fmt.Printf("Error while compiling %s: %v", src_file_name, err)
		return false
	}
	return true
}

func _link(project_config *ProjectConfig, build_dir_o_path, build_dir_out_path string) bool {
	o_files, err := os.ReadDir(build_dir_o_path)
	if err != nil {
		fmt.Printf("Error (reading object files): %v", err)
		return false
	}

	o := []string{}
	for _, o_file := range o_files {
		o = append(o, filepath.Join(build_dir_o_path, o_file.Name()))
	}

	if err := Link(
		project_config, // project config
		&o,             // object files
		filepath.Join(build_dir_out_path, project_config.Project.Name+".out"), // output path for executable
	); err != nil {
		fmt.Printf("Error (linking object files): %v", err)
		return false
	}
	return true
}
