package main

import (
	"os"
	"os/user"

	"github.com/salty-max/lars/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	repl.Start(os.Stdin, os.Stdout, user)
}
