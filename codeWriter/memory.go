package codewriter

import "strconv"

var memoryCommands = map[string]map[string]func(string, string) string{
	"C_PUSH": {
		"constant": pushConstant,
		"pointer":  pushPointer,
		"temp":     pushTemp,
		"argument": pushMemorySegment,
		"local":    pushMemorySegment,
		"static":   pushStatic,
		"this":     pushMemorySegment,
		"that":     pushMemorySegment,
	},
	"C_POP": {
		"constant": popConstant,
		"pointer":  popPointer,
		"temp":     popTemp,
		"argument": popMemorySegment,
		"local":    popMemorySegment,
		"static":   popStatic,
		"this":     popMemorySegment,
		"that":     popMemorySegment,
	},
}

// pop commands

func popStatic(fileName, index string) string {
	return "@SP\nM=M-1\nA=M\nD=M\n@" + fileName + "." + index + "\nM=D\n"
}

func popMemorySegment(segmentName string, index string) string {
	pointAtSegmentBaseAddress := "@" + assemblySegmentNotations[segmentName] + "\nA=M\n"
	segmentAddressOffsetByOne := "A=A+1\n"

	segmentAsInteger, _ := strconv.Atoi(index)

	if segmentAsInteger > 0 {
		for i := 0; i < segmentAsInteger; i++ {
			pointAtSegmentBaseAddress = pointAtSegmentBaseAddress + segmentAddressOffsetByOne
		}
	}
	return moveValueFromStackToDataRegister() + pointAtSegmentBaseAddress + "M=D\n"
}

func popPointer(segmentName string, index string) string {
	moveValueFromDataRegisterToPointerAddress := "@" + pointerSegments[index] + "\nM=D\n"
	return moveValueFromStackToDataRegister() + moveValueFromDataRegisterToPointerAddress
}

func popConstant(segmentName string, constantValue string) string {
	moveValueFromDataRegisterToConstantAddress := "@" + constantValue + "\nM=D\n"
	return moveValueFromStackToDataRegister() + moveValueFromDataRegisterToConstantAddress
}

// push commands

func pushStatic(fileName, index string) string {
	moveValueFromMemorySegmentToDataRegister := "@" + fileName + "." + index + "\nD=M\n"
	return moveValueFromMemorySegmentToDataRegister + moveValueFromDataRegisterToStack()
}

func pushConstant(segmentName string, constantValue string) string {
	moveValueFromConstantToDataRegister := "@" + constantValue + "\nD=A\n"
	return moveValueFromConstantToDataRegister + moveValueFromDataRegisterToStack()
}

func pushPointer(segmentName string, index string) string {
	moveValueFromMemorySegmentToDataRegister := "@" + pointerSegments[index] + "\nD=M\n"
	return moveValueFromMemorySegmentToDataRegister + moveValueFromDataRegisterToStack()
}

func popTemp(segmentName string, index string) string {
	tempBaseAddress := "@5\n"
	segmentAddressOffsetByOne := "A=A+1\n"

	segmentAsInteger, _ := strconv.Atoi(index)

	if segmentAsInteger > 0 {
		for i := 0; i < segmentAsInteger; i++ {
			tempBaseAddress += segmentAddressOffsetByOne
		}
	}

	return moveValueFromStackToDataRegister() + tempBaseAddress + "M=D\n"
}

func pushTemp(segmentName string, index string) string {
	tempBaseAddress := "@5\n"
	segmentAddressOffsetByOne := "A=A+1\n"

	segmentAsInteger, _ := strconv.Atoi(index)

	if segmentAsInteger > 0 {
		for i := 0; i < segmentAsInteger; i++ {
			tempBaseAddress = tempBaseAddress + segmentAddressOffsetByOne
		}
	}

	return tempBaseAddress + "D=M\n" + moveValueFromDataRegisterToStack()
}

func pushMemorySegment(segmentName string, segmentIndex string) string {
	pointAtSegmentBaseAddress := "@" + assemblySegmentNotations[segmentName] + "\nA=M\n"
	segmentAddressOffsetByOne := "A=A+1\n"

	segmentAsInteger, _ := strconv.Atoi(segmentIndex)

	if segmentAsInteger > 0 {
		for i := 0; i < segmentAsInteger; i++ {
			pointAtSegmentBaseAddress = pointAtSegmentBaseAddress + segmentAddressOffsetByOne
		}
	}
	return pointAtSegmentBaseAddress + "D=M\n" + moveValueFromDataRegisterToStack()
}

func moveValueFromDataRegisterToStack() string {
	return "@SP\nA=M\nM=D\n@SP\nM=M+1\n"
}

// stores stack value in D register and pops it out
func moveValueFromStackToDataRegister() string {
	return "@SP\nM=M-1\nA=M\nD=M\n"
}

var assemblySegmentNotations = map[string]string{
	"argument": "ARG",
	"local":    "LCL",
	"static":   "16",
	"this":     "THIS",
	"that":     "THAT",
	"temp":     "Temp",
}

var pointerSegments = map[string]string{
	"0": "THIS",
	"1": "THAT",
}
