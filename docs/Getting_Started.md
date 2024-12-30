# Getting Started

List of contents
* [Installation](#installation)
* [Initializing a Project](#initializing-the-project)
* [Adding an external library](#adding-an-external-library)
* [Building the Project]()
* [Running the executable]()

## Installation

Cherry is entirely written in go so it can be installed and build using the go.

> Make sure you have GO installed in your system, by running ``` go --version ``` command.<br>If its not installed. install it!<br>Also make sure the $GOBIN path Enviornment variable is set. check if it set or not by running this command: ``` echo $GOBIN ```

To install cherry, run this command:
```sh
go install github.com/yashtajne/cherry
```
Veryify the installation by running this command:
```sh
cherry version
```
<br>

## Initializing the Project

To create a C/C++ project, run this command:

```sh
cherry init <project-name>
```

This will create a cherry.toml project config file in the directory where the command is being executed.<br>It will also create src and build directories

<br>

## Adding an external library

Cherry relies on vcpkg to install the libraries.

> Make sure vcpkg is installed in your system if its not, install it!<br>Then also make sure that $VCPKG_ROOT path Enviornment varibale is set. check if its set or not by running this command: ``` echo $VCPKG_ROOT ```

To add an external library to your project<br>Install the library using vcpkg first by running this command:

```sh
vcpkg install <your-library-name>
```

After that run this command to add that library to your project directory.

```sh
cherry add <your-library-name>
```
This will add all the include files and static libraries to you project
<br>

## Make the Executable

To build the project, run this command:

```sh
cherry make
```
This will compile all the source files in you project then link them

## Run the Executable

To run the executable, run this command:

```sh
cherry run
```
This will execute the project binary in the terminal.