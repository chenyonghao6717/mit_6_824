package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print(("Go runs on "))
	switch os := runtime.GOOS; os {
		case "darmin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux.")
		default:
			fmt.Printf("%s.\n", os)
	}
}