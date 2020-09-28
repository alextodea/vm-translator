package codewriter

var arithmeticCommands = map[string]func() string{
	"add": arithmeticAdd,
	"sub": arithmeticSub,
	"neg": arithmeticNeg,
	"eq":  arithmeticEq,
	"gt":  arithmeticGt,
	"lt":  arithmeticLt,
	"and": arithmeticAnd,
	"or":  arithmeticOr,
	"not": arithmeticNot,
}

func arithmeticLabelGreatherThan() string {
	return "(COMPARISONISTRUE)\nM=-1\n(COMPARISONISNOTTRUE)\nM=0\n"
}

func arithmeticNot() string {
	return "@SP\nM=M-1\nA=M\nM=!M\n"
}

func arithmeticOr() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M|D\n"
}

func arithmeticAnd() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nM=M&D\n"
}

func arithmeticGt() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nD=M-D\n@COMPARISONISTRUE\nD;JGT\n@COMPARISONISNOTTRUE\n"
}

func arithmeticLt() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nD=M-D\n@COMPARISONISTRUE\nD;JLT\n@COMPARISONISNOTTRUE\n"
}

func arithmeticEq() string {
	return "@SP\nM=M-1\nA=M\nD=M\nA=A-1\nD=M-D\n@COMPARISONISTRUE\nD;JEQ\n@COMPARISONISNOTTRUE\n"
}

func arithmeticAdd() string {
	return "@SP\nA=M-1\nD=M\nA=A-1\nM=M+D\n@SP\nM=M-1\n"
}

func arithmeticSub() string {
	return "@SP\nA=M-1\nD=M\nA=A-1\nM=M-D\n@SP\nM=M-1\n"
}

func arithmeticNeg() string {
	return "@SP\nA=M-1\nM=-M\n"
}
