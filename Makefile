all:
	gofmt -s -w .
	go build
	
run:
	go run main.go
