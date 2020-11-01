

format:
	go fmt ./...
	golint ./...

build: format
	go build -o yamltojson