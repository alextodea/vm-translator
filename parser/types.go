package parser

// FileLineParser handles the parsing of a single .vm file, and encapsulates access to the input code.
// It reads VM commands, parses them, and provides convenient access to their components.
// In addition, it removes all white space and comments.
type FileLineParser interface {
	ParseVMFileLine() (parsedVMFileLine string, err error)
	// GetCommandType() (string, err error)
	// GetCommandArg1() (string, err error)
	// GetCommandArg2() (string, err error)
}

// FileLine represents a .vm file line 
type FileLine struct {
	textInput string
}
