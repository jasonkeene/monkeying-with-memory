
all: hello loop

hello: hello.o
	ld hello.o -o hello

hello.o:
	nasm -felf64 hello.s

loop: loop.o
	ld loop.o -o loop

loop.o:
	nasm -felf64 loop.s

docker-run: docker-build
	docker run -it --rm -v "$(PWD)":/pwd monkey-env

docker-build:
	docker build -t monkey-env .

clean:
	rm -f hello.o hello loop.o loop && docker rmi -f monkey-env
