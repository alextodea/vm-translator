package parser

// FileLine represents input file line and it expects a string
type FileLine string

// CleanedFileLine represents file line that does not contain white space, or is not a commented line
type CleanedFileLine string

// CommandType represents type of current VM command, derived from
type CommandType string

// FirstCommandArgument represents the first argument of the current command.
type FirstCommandArgument string

// SecondCommandArgument represents the second argument of the current command.
type SecondCommandArgument string

// Parser handles the parsing of a single .vm file, and encapsulates access to the input code.
// It reads VM commands, parses them, and provides convenient access to their components.
// In addition, it removes all white space and comments.
type Parser interface {
	cleanFileLine(fileLine FileLine) (cleanedFileLine CleanedFileLine, err error)
	getCommandType(cleanedFileLine CleanedFileLine) (commandType CommandType, err error)
	getCommandArg1(cleanedFileLine CleanedFileLine) (firstCommandArgument FirstCommandArgument, err error)
	getCommandArg2(cleanedFileLine CleanedFileLine) (secondCommandArgument SecondCommandArgument, err error)
}
