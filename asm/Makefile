
all: hello loop

hello: hello.o
	ld hello.o -o hello

hello.o:
	nasm -felf64 hello.s

loop: loop.o
	ld loop.o -o loop

loop.o:
	nasm -felf64 loop.s

clean:
	rm -f hello.o hello loop.o loop
