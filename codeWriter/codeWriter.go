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

	for _, vmCommand := range parsedCommands {
		commandType := vmCommand.CommandType
		firstCommandArg := vmCommand.Args[0]

		if commandType == "C_ARITHMETIC" {
			writeArithmetic(commandType, firstCommandArg)
		} else if commandType == "C_PUSH" || commandType == "C_POP" {
			secondCommandArg := vmCommand.Args[1]
			assemblySnippet, err := writePushPop(commandType, firstCommandArg, secondCommandArg)

			if err != nil {
				return err
			}

			outputFile.WriteString(assemblySnippet)
		}
	}
	outputFile.Close()
	return nil
}
