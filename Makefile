build:
	@go build -o bin/go-backend-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/go-backend-api