package codewriter

var memorySegments = map[string]string{
	"argument": "ARG",
	"local":    "LCL",
	"static":   "16",
	"this":     "THIS",
	"that":     "THAT",
	"temp": "Temp",
}

var pushFromDataRegisterToStackAssemblySnippet string = "@SP\nA=M\nM=D\n@SP\nM=M+1\n"

var pointerSegment = map[int]string {
	0: "@THIS\n",
	1: "@THAT\n",
}