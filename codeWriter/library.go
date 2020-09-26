package codewriter

import (
	"errors"
	"fmt"
	"strconv"
)

// Writes the assembly code that is the translation of the given arithmetic command.
func writeArithmetic(commandType string, firstCommandArg string) {
	fmt.Println("arithmetic", commandType, firstCommandArg)
}

func writePushPop(commandType string, firstCommandArg string, secondCommandArg string) (string, error) {
	assemblyBaseSnippet := "//" + commandType + ": " + firstCommandArg + " " + secondCommandArg + "\n"

	if commandType == "C_PUSH" {

		switch firstCommandArg {
		case "constant":
			pointToAddressPointedByConstant := "@" + secondCommandArg + "\n" // @i
			storeAddressNrToDataRegister := "D=A\n"                          // D= i
			pushConstantToStackAssemblySnippet := pointToAddressPointedByConstant + storeAddressNrToDataRegister + pushFromDataRegisterToStackAssemblySnippet
			assemblyBaseSnippet = assemblyBaseSnippet + pushConstantToStackAssemblySnippet
		case "temp":
			pointToTempAddress := "@" + memorySegments[firstCommandArg] + secondCommandArg + "\n" // @Temp+i
			storeSegmentValueToDataRegister := "D=M\n"                                            // D=M
			pushTempToStackAssemblySnippet := pointToTempAddress + storeSegmentValueToDataRegister + pushFromDataRegisterToStackAssemblySnippet
			assemblyBaseSnippet = assemblyBaseSnippet + pushTempToStackAssemblySnippet
		case "pointer":
			secondArgToInteger, err := strconv.Atoi(secondCommandArg)
			if err != nil {
				return "", errors.New("push pointer: could not convert command's 2nd argument to integer")
			}

			pointToThisOrThatSegments := pointerSegment[secondArgToInteger] // @THIS or @THAT
			storeSegmentValueToDataRegister := "D=M\n"                      // D=M
			pushPointerToStackAssemblySnippet := pointToThisOrThatSegments + storeSegmentValueToDataRegister + pushFromDataRegisterToStackAssemblySnippet
			assemblyBaseSnippet = assemblyBaseSnippet + pushPointerToStackAssemblySnippet
		default: // for all other memory segments
			pointToSegmentBasePointerAddress := "@" + firstCommandArg + "\n"
			pointToAddressOfSegmentAtSpecifiedIndex := "A=M+" + secondCommandArg + "\n"
			storeSegmentValueToDataRegister := "D=M\n" // D=M
			pushWordValueToStack := pointToSegmentBasePointerAddress + pointToAddressOfSegmentAtSpecifiedIndex + storeSegmentValueToDataRegister + pushFromDataRegisterToStackAssemblySnippet
			assemblyBaseSnippet = assemblyBaseSnippet + pushWordValueToStack
		}
	}

	return assemblyBaseSnippet, nil
}
