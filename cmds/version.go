package cmds

import (
	"fmt"
)

var PrintVersionMessage = `
Cherry v%s
Copyright (C) 2024-2025 Yash Tajne

This program may be freely redistributed
under the terms of the MIT License.
`

func Version(version string) {
	fmt.Printf(PrintVersionMessage, version)
}
