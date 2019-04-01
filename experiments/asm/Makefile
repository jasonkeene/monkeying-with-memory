
all: build/hello build/loop build/nested build/args build/call_c build/mult.o build/mult build/max_mult

build/hello: build/hello.o
	ld build/hello.o -o build/hello

build/hello.o: build
	nasm -f elf64 -o build/hello.o hello.s

build/loop: build/loop.o
	ld build/loop.o -o build/loop

build/loop.o: build
	nasm -f elf64 -o build/loop.o loop.s

build/nested: build/nested.o
	ld build/nested.o -o build/nested

build/nested.o: build
	nasm -f elf64 -o build/nested.o nested.s

build/args: build/args.o
	ld build/args.o -o build/args

build/args.o: build
	nasm -f elf64 -o build/args.o args.s

build/call_c: build/call_c.o
	ld build/call_c.o -o build/call_c -lc --dynamic-linker /lib/x86_64-linux-gnu/ld-2.24.so

build/call_c.o: build
	nasm -f elf64 -o build/call_c.o call_c.s

build/mult: build/mult.o
	gcc build/mult.o mult.c -o build/mult

build/mult.o: build
	nasm -f elf64 -o build/mult.o mult.s

build/max_mult: build/max_mult.o
	ld build/max_mult.o -o build/max_mult

build/max_mult.o: build
	nasm -f elf64 -o build/max_mult.o max_mult.s

build:
	mkdir build

docker-run: docker-build
	docker run -it --rm -v "$(PWD)":/pwd monkey-env

docker-build:
	docker build -t monkey-env .

docker-rmi:
	docker rmi -f monkey-env

clean:
	rm -rf build
