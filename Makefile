.PHONY: build
build:
	go build -o build/bin/web-server-go main.go

.PHONY: run
run:
	go run main.go