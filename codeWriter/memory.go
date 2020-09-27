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
	// 	"C_POP": {
	// 		"constant": pushConstantToStack,
	// 		"pointer":  pushConstantToStack,
	// 		"temp":     pushConstantToStack,
	// 		"argument": pushConstantToStack,
	// 		"local":    pushConstantToStack,
	// 		"static":   pushConstantToStack,
	// 		"this":     pushConstantToStack,
	// 		"that":     pushConstantToStack,
	// 	},
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

func pushCommandCommentedLine(firstCommandArg string, secondCommandArg string) string {
	return "//push" + firstCommandArg + secondCommandArg + "\n"
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
