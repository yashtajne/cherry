package cmds

import (
	"fmt"
)

func Version(version string) {
	fmt.Println("Cherry ", version)
	fmt.Println("Copyright (C) 2024 Yash Tajne")
	fmt.Println("This program may be freely redistributed")
	fmt.Println("under the terms of the MIT License.")
}
