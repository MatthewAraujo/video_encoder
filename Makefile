build:
	@go build -o bin/encoder cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/ecom