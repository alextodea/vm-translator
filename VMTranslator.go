package main

import (
	"fmt"
	"os"
	parser "vm-translator/parser"
)

func main() {
	err := parser.ParseVMFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("program finished")
}
