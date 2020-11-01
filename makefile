

format:
	go fmt ./...
	golint ./...

build: format
	go build -o yamltojson

install:
	go build -o yamltojson
	sudo mv yamltojson /usr/local/bin