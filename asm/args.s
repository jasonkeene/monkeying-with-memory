global _start

_start:
    push 21
    call times2

    push rax
    mov rdi, rax
    call exit

exit:
    mov rdi, [rsp+8]
    mov rax, 60
    syscall

times2:
    push rbp
    mov rbp, rsp

    mov rax, [rbp+16]
    add rax, rax

    mov rsp, rbp
    pop rbp
    ret
