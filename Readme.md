# Cherry

Cherry is a build tool for C and C++ projects that automates the compilation and linking process. It can also add vcpkg packages to to your projects.

## Features

- **Automated Compilation and Linking**: Simplify your build process with automated compilation and linking.
- **vcpkg Integration**: Add and manage vcpkg packages in your projects.
- **Cross-Platform Support**: Its entirely written in go so it will work on Windows, macOS, and Linux.

## How it works

Cherry compiles each source file separately in different processes. All compilation and linking commands are executed by the application itself, without relying on external build scripts.

It does not make use of MakeFiles and does not create any make files either.

It keeps track of file modifications in a log file (`cherry.log`). When a file is modified, It will automatically recompile only the changed files. All the compiled files will be cached.

#### Project Directory structure

```sh
├── include/
├── lib/
│   ├── release/
│   ├── debug/
│   └── pkgconfig/
├── build/
│   ├── release/
│   ├── debug/
│   └── cherry.log
├── src/
└── cherry.toml
```

### Configuration File

It used .toml as its configuration language
The config file ``` cherry.toml ``` contains
Project information, Build system information and list of packages

```toml
[Project]
  name = "project_name"
  description = "description of the project"
  version = "1.0"

[Build]
  os = "linux"
  shell = "/path/to/shell"
  compiler = "g++" # will be gcc for c projects
  includedir = "/path/to/include/directory"
  libdir = "/path/to/lib/directory"

[[Packages]]
  name = "package_name"
  description = "package_description"
  url = ""
  version = "package_version"
  libs = "\"-L${libdir}\" -libs" # libraries
  cflags = "\"-I${includedir}\"" # Compilation flags

```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.txt) file for details.
