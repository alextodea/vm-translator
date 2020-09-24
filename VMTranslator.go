package main

import (
	"fmt"
	"os"
	parser "vm-translator/parser"
)

func main() {
	parsedCommands, err := parser.Parser(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(parsedCommands)
	fmt.Println("program finished")
}
