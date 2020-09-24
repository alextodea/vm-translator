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
		var fileLineParser FileLineParser = FileLine{input.Text()}

		parsedVMFileLine, err := fileLineParser.ParseVMFileLine()

		if err != nil {
			return errors.New("failed to parse .vm file line")
		}

		fmt.Println(parsedVMFileLine)
	}

	return nil
}

// func cleanFileLine(fileLine string) Cl (cleanedFileLine string, err error) {
// 	isLineCommented := strings.HasPrefix(fileLine, "//")

// 	if len(fileLine) < 1 {
// 		return "", errors.New("file line is empty")
// 	}

// 	if isLineCommented {
// 		return "", errors.New("file line is commented")
// 	}

// 	return fileLine, nil
// }

// func getCommandType(cleanedFileLine string) (commandType string, err error) {
// 	return "", nil
// }
// func getCommandArg1(cleanedFileLine string) (firstCommandArgument string, err error) {
// 	return "", nil
// }
// func getCommandArg2(cleanedFileLine string) (secondCommandArgument string, err error) {
// 	return "", nil
// }
