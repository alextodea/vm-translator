// Package codewriter translates VM commands into Hack assembly code.
package codewriter

import (
	"os"
	parser "vm-translator/parser"
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

	for _, vmCommand := range parsedCommands {
		commandType := vmCommand.CommandType
		firstCommandArg := vmCommand.Args[0]

		switch commandType {
		case "C_PUSH":
			// 	fallthrough
			// case "C_POP":
			secondCommandArg := vmCommand.Args[1]

			assemblyCommand = memoryCommands[commandType][firstCommandArg](firstCommandArg, secondCommandArg)
		}
		outputFile.WriteString(assemblyCommand)
	}

	outputFile.Close()
	return nil
}
