// Package codewriter translates VM commands into Hack assembly code.
package codewriter

import (
	"os"

	parser "github.com/alextodea/vm-translator/parser"
)

// TranslateVMInstructionsToAssembly translates VM commands into Hack assembly code.
func TranslateVMInstructionsToAssembly(parsedCommands []parser.ParsedCommand, inputFileName string) (err error) {

	outputFileNameAndExtension := inputFileName + ".asm"
	outputFile, err := os.Create(outputFileNameAndExtension)

	if err != nil {
		outputFile.Close()
		return err
	}

	var assemblyCommand string
	var existsArithmeticLabelGreatherThan bool = false

	for _, vmCommand := range parsedCommands {
		commandType := vmCommand.CommandType
		firstCommandArg := vmCommand.Args[0]
		var baseAssemblySnippet string

		if !existsArithmeticLabelGreatherThan && firstCommandArg == "eq" || firstCommandArg == "lt" || firstCommandArg == "gt" {
			baseAssemblySnippet = baseAssemblySnippet + arithmeticLabelGreatherThan()
			existsArithmeticLabelGreatherThan = true
		}

		baseAssemblySnippet = baseAssemblySnippet + "//" + commandType + " " + firstCommandArg

		switch commandType {
		case "C_PUSH":
			fallthrough
		case "C_POP":
			secondCommandArg := vmCommand.Args[1]
			baseAssemblySnippet = baseAssemblySnippet + " " + secondCommandArg

			if firstCommandArg == "static" {
				assemblyCommand = memoryCommands[commandType][firstCommandArg](inputFileName, secondCommandArg)
			} else {
				assemblyCommand = memoryCommands[commandType][firstCommandArg](firstCommandArg, secondCommandArg)
			}

		case "C_ARITHMETIC":
			assemblyCommand = arithmeticCommands[firstCommandArg]()
		}

		baseAssemblySnippet = baseAssemblySnippet + "\n"
		baseAssemblySnippet = baseAssemblySnippet + assemblyCommand
		if firstCommandArg == "add" || firstCommandArg == "sub" || firstCommandArg == "neg" || firstCommandArg == "eq" {
			outputFile.WriteString(baseAssemblySnippet)
		}

	}

	outputFile.Close()
	return nil
}
