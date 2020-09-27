package codewriter

var memoryCommands = map[string]map[string]func(string, string) string{
	"C_PUSH": {
		"constant": pushConstantToStack,
		"pointer":  pushPointerToStack,
		"temp":     pushTempToStack,
		"argument": pushMemorySegmentValueToStack,
		"local":    pushMemorySegmentValueToStack,
		"static":   pushMemorySegmentValueToStack,
		"this":     pushMemorySegmentValueToStack,
		"that":     pushMemorySegmentValueToStack,
	},
	"C_POP": {
		"constant": popConstantFromStack,
		"pointer":  popPointerFromStack,
		"temp":     popTempFromStack,
		"argument": popMemorySegmentValueFromStack,
		"local":    popMemorySegmentValueFromStack,
		"static":   popMemorySegmentValueFromStack,
		"this":     popMemorySegmentValueFromStack,
		"that":     popMemorySegmentValueFromStack,
	},
}

func popMemorySegmentValueFromStack(segmentName string, segmentIndex string) string {
	commentedLine := popCommandCommentedLine(segmentName, segmentIndex)
	storeDataRegisterValueToSegment := "@" + assemblySegmentNotations[segmentName] + "\nA=M+" + segmentIndex + "\nM=D\n"
	return commentedLine + assemblySnippetPopValueFromStack() + storeDataRegisterValueToSegment
}

func popTempFromStack(firstCommandArg string, tempIndex string) string {
	commentedLine := popCommandCommentedLine(firstCommandArg, tempIndex)
	storeDataRegisterValueToTemp := "@Temp" + tempIndex + "\nM=D\n"
	return commentedLine + assemblySnippetPopValueFromStack() + storeDataRegisterValueToTemp
}

func popPointerFromStack(firstCommandArg string, pointerIndex string) string {
	commentedLine := popCommandCommentedLine(firstCommandArg, pointerIndex)
	storeDataRegisterValueToSegment := "@" + pointerSegments[pointerIndex] + "\nM=D\n"
	return commentedLine + assemblySnippetPopValueFromStack() + storeDataRegisterValueToSegment
}

func popConstantFromStack(firstCommandArg string, constantValue string) string {
	commentedLine := popCommandCommentedLine(firstCommandArg, constantValue)
	storeDataRegisterValueToConstantAddress := "@" + constantValue + "\nM=D\n"
	return commentedLine + assemblySnippetPopValueFromStack() + storeDataRegisterValueToConstantAddress
}

func pushConstantToStack(firstCommandArg string, constantValue string) string {
	commentedLine := pushCommandCommentedLine(firstCommandArg, constantValue)
	storeConstantValueToDataRegister := "@" + constantValue + "\nD=A\n"
	return commentedLine + storeConstantValueToDataRegister + assemblySnippetPushValueToStack()
}

func pushPointerToStack(firstCommandArg string, pointerIndex string) string {
	commentedLine := pushCommandCommentedLine(firstCommandArg, pointerIndex)
	storeSegmentValueToDataRegister := "@" + pointerSegments[pointerIndex] + "\nD=M\n"
	return commentedLine + storeSegmentValueToDataRegister + assemblySnippetPushValueToStack()
}

func pushTempToStack(firstCommandArg string, tempIndex string) string {
	commentedLine := pushCommandCommentedLine(firstCommandArg, tempIndex)
	storeTempValueToDataRegister := "@Temp" + tempIndex + "\nD=M\n"
	return commentedLine + storeTempValueToDataRegister + assemblySnippetPushValueToStack()
}

func pushMemorySegmentValueToStack(segmentName string, segmentIndex string) string {
	commentedLine := pushCommandCommentedLine(segmentName, segmentIndex)
	storeSegmentValueToDataRegister := "@" + assemblySegmentNotations[segmentName] + "\nA=M+" + segmentIndex + "\nD=M\n"
	return commentedLine + storeSegmentValueToDataRegister + assemblySnippetPushValueToStack()
}

func assemblySnippetPushValueToStack() string {
	return "@SP\nA=M\nM=D\n@SP\nM=M+1\n"
}

// stores stack value in D register and pops it out
func assemblySnippetPopValueFromStack() string {
	return "@SP\nA=M\nD=M\n@SP\nM=M-1\n"
}

func pushCommandCommentedLine(firstCommandArg string, secondCommandArg string) string {
	return "// push" + firstCommandArg + secondCommandArg + "\n"
}

func popCommandCommentedLine(firstCommandArg string, secondCommandArg string) string {
	return "// pop" + firstCommandArg + secondCommandArg + "\n"
}

var assemblySegmentNotations = map[string]string{
	"argument": "ARG",
	"local":    "LCL",
	"static":   "16",
	"this":     "THIS",
	"that":     "THAT",
}

var pointerSegments = map[string]string{
	"0": "THIS",
	"1": "THAT",
}
