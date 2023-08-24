build:
	@go build .

install-deps: 
	@go mod tidy
