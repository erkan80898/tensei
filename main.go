package main

import (
	"fmt"
	"os"
	"os/user"
	"tensei/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	println("Tensei V0.01")
	fmt.Printf("USER: %s\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
