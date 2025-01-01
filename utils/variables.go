package utils

var ProjectWorkDirectoryPath string
var ProjectSrcDirectoryPath string
var ProjectBuildDirectoryPath string
var ProjectIncludeDirectoryPath string
var ProjectLibDirectoryPath string
var ProjectConfigFilePath string
var ProjectLogFilePath string

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
`
