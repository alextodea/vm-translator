//C_PUSH constant 10
@10
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP local 0
@SP
M=M-1
A=M
D=M
@LCL
A=M
M=D
//C_PUSH constant 21
@21
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_PUSH constant 22
@22
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP argument 2
@SP
M=M-1
A=M
D=M
@ARG
A=M
A=A+1
A=A+1
M=D
//C_POP argument 1
@SP
M=M-1
A=M
D=M
@ARG
A=M
A=A+1
M=D
//C_PUSH constant 36
@36
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP this 6
@SP
M=M-1
A=M
D=M
@THIS
A=M
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
M=D
//C_PUSH constant 42
@42
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_PUSH constant 45
@45
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP that 5
@SP
M=M-1
A=M
D=M
@THAT
A=M
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
M=D
//C_POP that 2
@SP
M=M-1
A=M
D=M
@THAT
A=M
A=A+1
A=A+1
M=D
//C_PUSH constant 510
@510
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP temp 6
@SP
M=M-1
A=M
D=M
@Temp6
M=D
//C_PUSH local 0
@LCL
A=M
D=M
@SP
A=M
M=D
@SP
M=M+1
//C_PUSH that 5
@THAT
A=M
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
D=M
@SP
A=M
M=D
@SP
M=M+1
//C_ARITHMETIC add
@SP
A=M-1
D=M
A=A-1
M=M+D
@SP
M=M-1
//C_PUSH argument 1
@ARG
A=M
A=A+1
D=M
@SP
A=M
M=D
@SP
M=M+1
//C_ARITHMETIC sub
@SP
A=M-1
D=M
A=A-1
M=M-D
@SP
M=M-1
//C_PUSH this 6
@THIS
A=M
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
D=M
@SP
A=M
M=D
@SP
M=M+1
//C_PUSH this 6
@THIS
A=M
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
D=M
@SP
A=M
M=D
@SP
M=M+1
//C_ARITHMETIC add
@SP
A=M-1
D=M
A=A-1
M=M+D
@SP
M=M-1
//C_ARITHMETIC sub
@SP
A=M-1
D=M
A=A-1
M=M-D
@SP
M=M-1
//C_PUSH temp 6
@Temp0
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
A=A+1
D=M
@SP
A=M
M=D
@SP
M=M+1
//C_ARITHMETIC add
@SP
A=M-1
D=M
A=A-1
M=M+D
@SP
M=M-1
