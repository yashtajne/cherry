package utils

// error constants
const (
	PACKAGE_NOT_INSTALLED_LOCALLY = "package_not_installed_locally"
	NOT_A_PACKAGE                 = "not_a_package"
)

var Project_Work_Directory_Path string
var Project_Src_Directory_Path string

var Project_Build_Directory_Path string
var Project_Build_Release_Directory_Path string
var Project_Build_Debug_Directory_Path string

var Project_Include_Directory_Path string

var Project_Lib_Directory_Path string
var Project_Lib_Release_Directory_Path string
var Project_Lib_Debug_Directory_Path string

var Project_Config_File_Path string

var Project_Log_File_Path string

var DefaultMainCPPFile string = `
#include <iostream>

int main(int argc, char** argv) {
	std::cout << "Hello, World!" << std::endl;
	return 0;
}
`

var DefaultMainCFile string = `
#include<stdio.h>

int main(int argc, char** argv) {
	printf("Hello, World!");
	return 0;
}
`

var DefaultCommandGitignoreFile = `
build/*
include/*
lib/*
`
