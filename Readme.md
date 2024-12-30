# Cherry

Cherry is a build tool for C and C++ projects that automates the compilation and linking process. It can also add vcpkg packages to to your projects.

## Features

- **Automated Compilation and Linking**: Simplify your build process with automated compilation and linking.
- **vcpkg Integration**: Add and manage vcpkg packages in your projects.
- **Cross-Platform Support**: Its entirely written in go so it will work on Windows, macOS, and Linux.

## Installation

To install Cherry run this command  
you must have go (golang) installed in your system

```sh
go install github.com/yashtajne/cherry
```

Then verify the installation by running
make sure $GOBIN path is set

```sh
cherry version
```

### Initializing a project

To create a C/C++ project, run this command:

```sh
cherry init <your-project-name>
```

#### Project Directory structure

```sh
├── include/
├── lib/
├── build/
│   ├── o/
│   └── out/
├── src/
├── cherry.toml
└── cherry.log
```

### Adding vcpkg Packages

To add a vcpkg package to your project, use the following command:

```sh
cherry add <package-name>
```

You must have vcpkg installed and the package installed in your system locally.

### Build executable

To build your project, run this command:

```sh
cherry make
```

Then run your executable using this command:

```sh
cherry run
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
