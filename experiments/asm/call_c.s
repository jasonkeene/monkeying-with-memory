global _start
extern printf

section .text

_start:
    mov rsi, 1234
    mov rdi, msg
    mov rax, 0
    call printf

    mov rax, 60
    xor rdi, rdi
    syscall

section .data

msg:
    db "Calling from ASM into C: %d", 10, 0
