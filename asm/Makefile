
all: hello loop nested

hello: hello.o
	ld hello.o -o hello

hello.o:
	nasm -felf64 hello.s

loop: loop.o
	ld loop.o -o loop

loop.o:
	nasm -felf64 loop.s

nested: nested.o
	ld nested.o -o nested

nested.o:
	nasm -felf64 nested.s

docker-run: docker-build
	docker run -it --rm -v "$(PWD)":/pwd monkey-env

docker-build:
	docker build -t monkey-env .

clean:
	rm -f hello.o hello loop.o loop nested nested.o && docker rmi -f monkey-env
