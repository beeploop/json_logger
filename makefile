build:
	@go build -o bin/json_tail main.go

run: 
	@go run main.go

test:
	@go test -v ./...

clean:
	@rm -rf bin
