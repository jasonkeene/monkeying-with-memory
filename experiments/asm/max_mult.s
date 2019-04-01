
global _start

section .text

_start:
    mov rax, [max_int_64]
    mul rax
    mov [result_64_lower], rax
    mov [result_64_higher], rdx

    mov     rax, 1
    mov     rdi, 1
    mov     rsi, result_64_lower
    mov     rdx, 16
    syscall

    mov eax, [max_int_32]
    mul eax
    mov [result_32_lower], eax
    mov [result_32_higher], edx

    mov     rax, 1
    mov     rdi, 1
    mov     rsi, result_32_lower
    mov     rdx, 8
    syscall

    mov     rax, 60
    xor     rdi, rdi
    syscall

section .data

max_int_64:
    dq      0xFFFFFFFFFFFFFFFF
result_64_lower:
    dq      0
result_64_higher:
    dq      0
max_int_32:
    dd      0xFFFFFFFF
result_32_lower:
    dd      0
result_32_higher:
    dd      0
