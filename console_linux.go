// +build linux

package silo

import (
	"fmt"
)

// PrintConsole attaches a console and prints to it.
func PrintConsole(m string) {
	fmt.Println(m)
}
