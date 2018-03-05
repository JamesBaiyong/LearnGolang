package main

import (
	"fmt"
	"runtime"
)

func os_know() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.",os)
	}
}

func main() {
	fmt.Print("Go runs on ")
	os_know()
}
