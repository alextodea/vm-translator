package main

import (
	"fmt"
	"os"
	codeWriter "vm-translator/codeWriter"
	parser "vm-translator/parser"
)

func main() {
	inputFileName := os.Args[1]
	parsedCommands, err := parser.Parser(inputFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = codeWriter.TranslateVMInstructionsToAssembly(parsedCommands, inputFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("program finished")
}
