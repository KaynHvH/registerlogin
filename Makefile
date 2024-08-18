build:
	@go build -o bin/registerlogin

run: build
	@./bin/registerlogin

test:
	@go test -v ./...