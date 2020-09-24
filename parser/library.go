package parser

import (
	"errors"
	"strings"
)

// ParseVMFileLine handles the parsing of a single .vm file, and encapsulates access to the input code.
// It reads VM commands, parses them, and provides convenient access to their components.
func (fileLine FileLine) ParseVMFileLine() (parsedVMFileLine string, err error) {
	isLineCommented := strings.HasPrefix(fileLine.textInput, "//")

	if len(fileLine.textInput) < 1 {
		return "", errors.New("file line is empty")
	}

	if isLineCommented {
		return "", errors.New("file line is commented")
	}

	return fileLine.textInput, nil
}
