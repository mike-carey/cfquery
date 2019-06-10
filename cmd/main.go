package main

import (
	"os"

	"github.com/mike-carey/cfquery/commands"
)

func main() {
	err := commands.Main(os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	if err != nil {
		panic(err)
		os.Exit(1)
	}

	os.Exit(0)
}
