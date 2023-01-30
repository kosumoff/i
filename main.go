package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	// Get the current user.
	cu, err := user.Current()
	if err != nil {
		return
	}

	// Check if the program was invoked with 2 arguments.
	if len(os.Args) != 3 {
		return
	}

	// Check if the arguments are "love" and "you".
	if strings.EqualFold(os.Args[1], "love") && strings.EqualFold(os.Args[2], "you") {
		if cu.Username != "root" {
			fmt.Println("I have a boyfriend")
		} else {
			fmt.Println("I love you too")
		}
	}
}
