

format:
	go fmt ./...
	golint ./...

build: format
	go build -o yamltojson

zip:
	go build -o yamltojson
	chmod +x yamltojson
	zip -9 yamltojson.zip yamltojson
	
install:
	go build -o yamltojson
	chmod +x yamltojson
	sudo mv yamltojson /usr/local/bin