// 0 C_PUSH constant 111
(STATICTEST0)
@111
D=A
@SP
A=M
M=D
@SP
M=M+1
// 1 C_PUSH constant 333
(STATICTEST1)
@333
D=A
@SP
A=M
M=D
@SP
M=M+1
// 2 C_PUSH constant 888
(STATICTEST2)
@888
D=A
@SP
A=M
M=D
@SP
M=M+1
// 3 C_POP static 8
(STATICTEST3)
@SP
M=M-1
A=M
D=M
@static.8
M=D
// 4 C_POP static 3
(STATICTEST4)
@SP
M=M-1
A=M
D=M
@static.3
M=D
// 5 C_POP static 1
(STATICTEST5)
@SP
M=M-1
A=M
D=M
@static.1
M=D
// 6 C_PUSH static 3
(STATICTEST6)
@static.3
D=M
@SP
A=M
M=D
@SP
M=M+1
// 7 C_PUSH static 1
(STATICTEST7)
@static.1
D=M
@SP
A=M
M=D
@SP
M=M+1
// 8 C_ARITHMETIC sub
(STATICTEST8)
@SP
A=M-1
D=M
A=A-1
M=M-D
@SP
M=M-1
// 9 C_PUSH static 8
(STATICTEST9)
@static.8
D=M
@SP
A=M
M=D
@SP
M=M+1
// 10 C_ARITHMETIC add
(STATICTEST10)
@SP
A=M-1
D=M
A=A-1
M=M+D
@SP
M=M-1
(END)
@END
0;JMP
