package codewriter

import (
	"strconv"
	"strings"
)

var arithmeticOperationCommands = map[string]func() string{
	"add": arithmeticAdd,
	"sub": arithmeticSub,
	"neg": arithmeticNeg,
	"and": arithmeticAnd,
	"or":  arithmeticOr,
	"not": arithmeticNot,
}

var arithmeticComparisonCommands = map[string]func(numberOfCommandsLeft, commandCounter int, inputFileName, firstCommandArg, commandCounterString string) string{
	"eq": handleComparisonCommands,
	"gt": handleComparisonCommands,
	"lt": handleComparisonCommands,
}

func handleComparisonCommands(numberOfCommandsLeft, commandCounter int, inputFileName, firstCommandArg, commandCounterString string) string {
	inputFileNameToUpperLetters := strings.ToUpper(inputFileName)
	assemblyCommand := "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nD=M-D\n"
	var jmp string

	switch firstCommandArg {
	case "eq":
		jmp = "D;JEQ\n"
	case "gt":
		jmp = "D;JGT\n"
	case "lt":
		jmp = "D;JLT\n"
	}

	assemblyComparisonBody := addComparisonBody(inputFileNameToUpperLetters, commandCounterString, jmp, numberOfCommandsLeft)
	assemblyCommand += assemblyComparisonBody
	return assemblyCommand
}

func addComparisonBody(inputFileNameToUpperLetters, commandCounterString, jmp string, numberOfCommandsLeft int) string {
	var body string
	atLabelComparisonIsTrue := "@" + inputFileNameToUpperLetters + "_COMPARISON_IS_TRUE_" + commandCounterString + "\n"
	atLabelComparisonIsNotTrue := "@" + inputFileNameToUpperLetters + "_COMPARISON_IS_NOT_TRUE_" + commandCounterString + "\n"
	body += atLabelComparisonIsTrue + jmp + atLabelComparisonIsNotTrue + "0;JMP\n"
	comparisonIsTrueLabel := addComparisonLabelBody(inputFileNameToUpperLetters, commandCounterString, "TRUE", "-1")
	comparisonIsNotTrueLabel := addComparisonLabelBody(inputFileNameToUpperLetters, commandCounterString, "NOT_TRUE", "0")

	var branchToNextCommand string
	if numberOfCommandsLeft == 0 {
		branchToNextCommand = "@END\n0;JMP\n"
	} else {
		commandCounter, _ := strconv.Atoi(commandCounterString)
		commandCounter++
		counterForNextCommandLabel := strconv.Itoa(commandCounter)
		branchToNextCommand = "@" + inputFileNameToUpperLetters + counterForNextCommandLabel + "\n0;JMP\n"
	}

	body += comparisonIsTrueLabel + branchToNextCommand + comparisonIsNotTrueLabel + branchToNextCommand
	return body
}

func addComparisonLabelBody(inputFileNameToUpperLetters, commandCounterString, trueNotTrue, registerValue string) string {
	body := "(" + inputFileNameToUpperLetters + "_COMPARISON_IS_" + trueNotTrue + "_" + commandCounterString + ")\n@SP\nA=M-1\nM=" + registerValue + "\n"
	return body
}

func arithmeticNot() string {
	return "@SP\nA=M-1\nM=!M\n"
}

func arithmeticOr() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M|D\n"
}

func arithmeticAnd() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M&D\n"
}

func arithmeticGt() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nD=M-D\n"
}

func arithmeticLt() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nD=M-D\n"
}

func arithmeticEq() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nD=M-D\n"

}

func arithmeticAdd() string {
	return "@SP\nA=M-1\nD=M\nA=A-1\nM=M+D\n@SP\nM=M-1\n"
}

func arithmeticSub() string {
	return "@SP\nA=M-1\nD=M\nA=A-1\nM=M-D\n@SP\nM=M-1\n"
}

func arithmeticNeg() string {
	return "@SP\nA=M-1\nM=-M\n"
}
