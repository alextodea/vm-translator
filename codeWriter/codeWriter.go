// Package codewriter translates VM commands into Hack assembly code.
package codewriter

import (
	"os"
	"path/filepath"
	"strings"

	"strconv"

	parser "github.com/alextodea/vm-translator/parser"
)

// TranslateVMInstructionsToAssembly translates VM commands into Hack assembly code.
func TranslateVMInstructionsToAssembly(parsedCommands []parser.ParsedCommand, inputFilePath string) (err error) {
	targetFilePathWithoutExtension := strings.Split(inputFilePath, ".")[0]
	targetFilePath := targetFilePathWithoutExtension + ".asm"

	if err := os.MkdirAll(filepath.Dir(targetFilePath), 0770); err != nil {
		return err
	}

	outputFile, err := os.Create(targetFilePath)
	inputFileName := parser.GetFileName(inputFilePath)

	numberOfCommandsLeft := len(parsedCommands)

	for commandCounter, vmCommand := range parsedCommands {
		numberOfCommandsLeft--
		commandType := vmCommand.CommandType
		firstCommandArg := vmCommand.Args[0]

		// assembly command base
		var assemblyCommand string
		commandCounterString := strconv.Itoa(commandCounter)

		// first line of assembly command
		// it is a commented line that denotes the command in vm language, as well as command counter
		assemblyCommandCommentedLine := addAssemblyCommandCommentedLine(vmCommand, commandCounterString)
		assemblyCommand += assemblyCommandCommentedLine

		// assembly command label declaration (XXX)
		assemblyCommandLabelDeclaration := addAssemblyCommandLabelDeclaration(inputFileName, commandCounterString)
		assemblyCommand += assemblyCommandLabelDeclaration

		if commandType == "C_ARITHMETIC" {
			// if comparison command
			if firstCommandArg == "eq" || firstCommandArg == "lt" || firstCommandArg == "gt" {
				assemblyCommand += arithmeticComparisonCommands[firstCommandArg](numberOfCommandsLeft, commandCounter, inputFileName, firstCommandArg, commandCounterString)
			} else {
				assemblyCommand += arithmeticOperationCommands[firstCommandArg]()
			}
		} else if commandType == "C_PUSH" || commandType == "C_POP" {
			if firstCommandArg == "STATIC" {
				assemblyCommand += memoryCommands[commandType][firstCommandArg](inputFileName, vmCommand.Args[1])
			} else {
				assemblyCommand += memoryCommands[commandType][firstCommandArg](firstCommandArg, vmCommand.Args[1])
			}
		}

		if numberOfCommandsLeft == 0 {
			assemblyCommand += "(END)\n@END\n0;JMP\n"
		}

		outputFile.WriteString(assemblyCommand)

	}

	outputFile.Close()
	return nil
}

func addAssemblyCommandLabelDeclaration(fileName, commandCounterString string) string {
	inputFileNameToUpperLetters := strings.ToUpper(fileName)
	assemblyCommandLabelDeclaration := "\n(" + inputFileNameToUpperLetters + commandCounterString + ")" + "\n"
	return assemblyCommandLabelDeclaration
}

func addAssemblyCommandCommentedLine(vmCommand parser.ParsedCommand, commandCounterString string) string {
	assemblyCommandCommentedLine := "//" + " " + commandCounterString + " " + vmCommand.CommandType

	for _, vmCommandArg := range vmCommand.Args {
		assemblyCommandCommentedLine += " " + vmCommandArg
	}

	return assemblyCommandCommentedLine
}
