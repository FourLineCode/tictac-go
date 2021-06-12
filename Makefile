all:
	go run src/*.go
build:
	go build -o bin/main src/*.go
run:
	./bin/main