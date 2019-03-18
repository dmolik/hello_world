
all: build push

build-code:
	GOOS=linux ARCH=amd64 go build -v -o hello_world ./

build-container:
	docker build -t graytshirt/hello_world:latest .

push:
	docker push graytshirt/hello_world:latest

build: build-code build-container

clean:
	rm -f hello_world

.PHONY: build build-code build-container push
