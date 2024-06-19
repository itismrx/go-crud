build:
	@~/go/bin/air  

run: build
	@./bin/jwt

test:
	@go test -v ./..@