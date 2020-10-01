//C_PUSH constant 3030
@3030
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP pointer 0
@SP
M=M-1
A=M
D=M
@THIS
M=D
//C_PUSH constant 3040
@3040
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP pointer 1
@SP
M=M-1
A=M
D=M
@THAT
M=D
//C_PUSH constant 32
@32
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP this 2
@SP
M=M-1
A=M
D=M
@THIS
A=M
A=A+1
A=A+1
M=D
//C_PUSH constant 46
@46
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP that 6
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
A=A+1
M=D
//C_PUSH pointer 0
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
//C_PUSH pointer 1
@THAT
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
//C_PUSH this 2
@THIS
A=M
A=A+1
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
//C_PUSH that 6
@THAT
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
