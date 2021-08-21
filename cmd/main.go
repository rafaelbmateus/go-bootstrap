package main

import (
	"fmt"
)

var (
	Version = "No version provided"
	Commit  = "No commit hash provided"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("I'm running with version: " + Version)
	fmt.Println("and signed with commit: " + Commit)
}
