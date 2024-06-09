package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/salty-max/lars/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Lars!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
