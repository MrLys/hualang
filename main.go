package main

import (
	"fmt"
	"os"
	"os/user"

	"ljos.app/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! This is the X progamming language", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}
