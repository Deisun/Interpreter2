package main

import (
	"os/user"
	"fmt"
	"Interpreter2/repl"
	"os"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!  This is the Monkey programming language!\n", user.Name)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
