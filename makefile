build:
	@go build -o bin/json_logger main.go logger.go

run: 
	@go run main.go

test:
	@go test -v ./...

clean:
	@rm -rf bin
