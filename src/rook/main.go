package main

import (
	"fmt"
	// "os"
	// "os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome to Rook.   >-/>\n")
}
