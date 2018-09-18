global _start

section .text

_start:
    xor     rax, rax        ; zero out rax
    mov     rcx, 33         ; how many times to increment
    call    inc
    call    print
    call    exit

inc:                        ; loop incrementing rax, rcx times
    inc     rax
    loop    inc
    ret

print:                      ; print the value of rax
    mov     rsi, message
    mov     [rsi], rax
    mov     rcx, 10         ; add newline character
    mov     [rsi+1], rcx

    mov     rax, 1
    mov     rdi, 1
    mov     rdx, 2
    syscall
    ret

exit:                       ; exit with 0
    mov     rax, 60
    xor     rdi, rdi
    syscall
    ret

section .data

message:
    db      0, 0            ; storage for output
