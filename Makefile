build:
	@go build -o bin/PiggyVault

run: build
	@./bin/PiggyVault
test:
	@go test -v ./...