package parser

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
)

// ParseVMFile parses a vm file, cleans its lines (remove whitespace, comments) and stores commands and arguments
func ParseVMFile(fileName string) (err error) {
	folderPath := "./vm-files/"
	filePath := path.Join(folderPath, fileName+".vm")

	f, err := os.Open(filePath)

	if err != nil {
		return errors.New("error opening file: " + fileName)
	}

	input := bufio.NewScanner(f)

	for input.Scan() {
		fileLine := input.Text()
		// val,err := cleanFileLine(fileLine)
		fmt.Println(fileLine)
	}

	return nil
}

func cleanFileLine(fileLine string) (cleanedFileLine CleanedFileLine, err error) {
	return "", nil
}

func getCommandType(cleanedFileLine CleanedFileLine) (commandType CommandType, err error) {
	return "", nil
}
func getCommandArg1(cleanedFileLine CleanedFileLine) (firstCommandArgument FirstCommandArgument, err error) {
	return "", nil
}
func getCommandArg2(cleanedFileLine CleanedFileLine) (secondCommandArgument SecondCommandArgument, err error) {
	return "", nil
}
