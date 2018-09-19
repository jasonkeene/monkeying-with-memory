global mult5

section .text

mult5:
    push rbp
    mov rbp, rsp

    mov rax, rdi ; copy first arg into rax
    mov rcx, 5
    mul rcx        ; multiply it by 5

    mov rsp, rbp
    pop rbp
    ret
