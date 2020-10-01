package main

import (
	"fmt"
	"os"

	codeWriter "github.com/alextodea/vm-translator/codeWriter"
	parser "github.com/alextodea/vm-translator/parser"
)

func main() {
	parsedCommands, err := parser.Parser(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inputFileName := parser.GetFileName(os.Args[1])
	err = codeWriter.TranslateVMInstructionsToAssembly(parsedCommands, inputFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("program finished")
}
