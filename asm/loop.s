		global	_start

		section	.text
_start:		xor	rax, rax
		mov	rcx, 33

loop:
		inc	rax
		loop	loop

		mov	rsi, message
		mov	[rsi], rax
		mov	rcx, 10
		mov	[rsi+1], rcx

		mov	rax, 1
		mov	rdi, 1
		mov	rdx, 2
		syscall

		mov 	rax, 60		; exit with 0
		xor	rdi, rdi
		syscall

		section	.data
message:	db	0, 0
