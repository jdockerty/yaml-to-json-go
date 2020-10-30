

format:
	go fmt ./...

build: format
	go build -o yamltojson