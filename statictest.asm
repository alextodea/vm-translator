//C_PUSH constant 111
@111
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_PUSH constant 333
@333
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_PUSH constant 888
@888
D=A
@SP
A=M
M=D
@SP
M=M+1
//C_POP static 8
@SP
M=M-1
A=M
D=M
@statictest.8
M=D
//C_POP static 3
@SP
M=M-1
A=M
D=M
@statictest.3
M=D
//C_POP static 1
@SP
M=M-1
A=M
D=M
@statictest.1
M=D
//C_PUSH static 3
@statictest.3
D=M
@SP
A=M
M=D
@SP
M=M+1
//C_PUSH static 1
@statictest.1
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
//C_PUSH static 8
@statictest.8
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
