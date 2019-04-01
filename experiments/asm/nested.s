global _start

section .text

_start:
    xor     rax, rax        ; zero out rax
    call    inc10
    call    print
    call    exit

inc10:
    call    inc5
    call    inc5
    ret

inc5:
    push rbp                ; store base pointer of previous func on stack
    mov rbp, rsp            ; store stack pointer as base pointer
    sub rsp, 1              ; allocate space on stack for data
    mov [rsp], byte 1       ; add data to stack
    add rax, [rsp]          ; use stack data to increment rax
    call inc4
    mov rsp, rbp            ; restore stack pointer from base pointer
    pop rbp                 ; restore base pointer of previous func
    ret

inc4:
    sub rsp, 8              ; manually store base pointer of previous func
    mov [rsp], qword rbp

    mov rbp, rsp            ; store stack pointer as base pointer
    sub rsp, 1              ; allocate space on stack for data
    mov [rsp], byte 4       ; add data to stack
    add rax, [rsp]          ; use stack data to increment rax
    mov rsp, rbp            ; restore stack pointer from base pointer

    mov rbp, [rsp]          ; manually restore base pointer of previous func
    add rsp, 8
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
