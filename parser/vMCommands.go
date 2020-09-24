package parser

var vmCommandsTable = map[string]string{
	"push": "C_PUSH",
	"pop":  "C_POP",
	"add":  "C_ARITHMETIC",
	"sub":  "C_ARITHMETIC",
	"neg":  "C_ARITHMETIC",
	"eq":   "C_ARITHMETIC",
	"gt":   "C_ARITHMETIC",
	"lt":   "C_ARITHMETIC",
	"and":  "C_ARITHMETIC",
	"or":   "C_ARITHMETIC",
	"not":  "C_ARITHMETIC",
}
