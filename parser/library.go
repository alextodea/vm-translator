package parser

import (
	"errors"
	"strings"
)

// ParseVMFileLine handles the parsing of a single .vm file, and encapsulates access to the input code.
// It reads VM commands, parses them, and provides convenient access to their components.
func parseFileLine(fileLine string) (parsedComponents []string, err error) {
	isLineCommented := strings.HasPrefix(fileLine, "//")

	if len(fileLine) < 1 {
		return parsedComponents, errors.New("file line is empty")
	}

	if isLineCommented {
		return parsedComponents, errors.New("file line is commented")
	}

	parsedComponents = strings.Split(fileLine, " ")

	if len(parsedComponents) < 1 {
		return parsedComponents, errors.New("parsed components array is empy")
	}

	return parsedComponents, nil
}

func getCommandFirstArgument(currentCommandType string, parsedComponents []string) (firstCommandArg string, err error) {

	if currentCommandType == "C_ARITHMETIC" {
		return parsedComponents[0], nil
	}

	if len(parsedComponents) < 2 {
		return "", errors.New("parsed components does not contain a second argument")
	}

	return parsedComponents[1], nil

}

func getCommandSecondArgument(currentCommandType string, parsedComponents []string) (secondCommandArg string, err error) {

	if currentCommandType == "C_PUSH" || currentCommandType == "C_POP" || currentCommandType == "C_FUNCTION" || currentCommandType == "C_CALL" {
		return parsedComponents[2], nil
	}

	return "", errors.New("command does not have a 2nd argument")
}
