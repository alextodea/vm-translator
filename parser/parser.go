package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
)

// ParsedCommand represents decomposed commands and their argument(s), as extracted from input file lines
type ParsedCommand struct {
	commandType string
	args        []string
}

// Parser parses a vm file, cleans its lines (remove whitespace, comments) and stores commands and arguments
func Parser(fileName string) (parsedCommands []ParsedCommand, err error) {
	folderPath := "./vm-files/"
	filePath := path.Join(folderPath, fileName+".vm")

	f, err := os.Open(filePath)

	if err != nil {
		return parsedCommands, errors.New("error opening file: " + fileName)
	}

	input := bufio.NewScanner(f)

	for input.Scan() {

		fileLine := input.Text()

		parsedComponents, err := parseFileLine(fileLine)

		if err != nil {
			fmt.Println(err)
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
			parsedCommands = append(parsedCommands, ParsedCommand{commandType: currentCommandType, args: []string{currentCommandFirstArgument}})
		} else {
			// append command type name and two args
			parsedCommands = append(parsedCommands, ParsedCommand{commandType: currentCommandType, args: []string{currentCommandFirstArgument, currentCommandSecondArgument}})
		}

	}

	return parsedCommands, nil
}
