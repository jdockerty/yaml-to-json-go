

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
	wget https://yaml-to-json-go.s3.eu-west-2.amazonaws.com/yamltojson.zip
	unzip yamltojson.zip
	sudo mv yamltojson /usr/local/bin
	rm yamltojson.zip