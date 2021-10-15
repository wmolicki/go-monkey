build:
	go build -o bin/monke.exe
test:
	go test -race ./...
format:
	gofmt -w .
