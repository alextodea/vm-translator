package parser

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// ParsedCommand represents decomposed commands and their argument(s), as extracted from input file lines
type ParsedCommand struct {
	CommandType string
	Args        []string
}

// Parser parses a vm file, cleans its lines (remove whitespace, comments) and stores commands and arguments
func Parser(filePath string) (parsedCommands []ParsedCommand, err error) {

	fileName := GetFileName(filePath)

	f, err := os.Open(filePath)

	if err != nil {
		return parsedCommands, errors.New("error opening file: " + fileName)
	}

	input := bufio.NewScanner(f)

	for input.Scan() {

		fileLine := input.Text()

		parsedComponents, err := parseFileLine(fileLine)

		if err != nil {
			continue
		}

		firstElementFromParsedComponents := parsedComponents[0]
		currentCommandType := vmCommandsTable[firstElementFromParsedComponents]

		currentCommandFirstArgument, err := getCommandFirstArgument(currentCommandType, parsedComponents)

		if err != nil {
			return parsedCommands, err
		}

		currentCommandSecondArgument, err := getCommandSecondArgument(currentCommandType, parsedComponents)

		if err != nil {
			// if no 2nd arguments exists, append only command type name and first arg
			parsedCommands = append(parsedCommands, ParsedCommand{CommandType: currentCommandType, Args: []string{currentCommandFirstArgument}})
		} else {
			// append command type name and two args
			parsedCommands = append(parsedCommands, ParsedCommand{CommandType: currentCommandType, Args: []string{currentCommandFirstArgument, currentCommandSecondArgument}})
		}

	}

	return parsedCommands, nil
}

// GetFileName extracts file name from path
func GetFileName(filePath string) string {
	splitInputPath := strings.Split(filePath, "/")

	inputFileNameAndExtension := splitInputPath[len(splitInputPath)-1]

	return strings.Split(inputFileNameAndExtension, ".")[0]
}
