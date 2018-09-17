
main:
	go build -gcflags '-N -l' -o main

dis: main
	objdump -M intel -d main

docker-build:
	docker build -t monkey-env .

docker-run: docker-build
	docker run -it --rm -v "$(PWD)":/pwd monkey-env

clean:
	rm -f main && docker rmi -f monkey-env
