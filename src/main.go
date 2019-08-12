package main

import (
	"fmt"
	"os"
	"os/user"
	"repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the monkey programing language!\n", user.Username)
	fmt.Printf("Feel to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
