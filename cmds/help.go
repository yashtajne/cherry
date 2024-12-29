package cmds

import "fmt"

func printHelpMessage() {
	fmt.Println("Usage: cherry <command> [arguments]")
	fmt.Println("operations:")
	fmt.Println("  cherry init    <project_name>   Initialize a new project")
	fmt.Println("  cherry add     <package_name>   Add a package to the project")
	fmt.Println("  cherry remove  <package_name>   Remove a package from the project")
	fmt.Println("  cherry make                     Build the project executable")
	fmt.Println("  cherry run                      Run the project executable")
	fmt.Println("  cherry help                     Display help and list commands")
	fmt.Println("  cherry version                  Display version information")
}

func Help(sub_cmd string) {
	switch sub_cmd {
	default:
		printHelpMessage()
	}
}
