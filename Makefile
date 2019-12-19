.DEFAULT_GOAL := help

help:
	@echo "Available options:\n- build: Build new version\n- test: Run all tests\n- clean: Clean project\n- check: Test and check project\n- imports: run goimports with golangci rules"

build: clean
	@printf "%s" "Building darwin/amd64..."
	@env GOOS=darwin GOARCH=amd64 go build -o builds/pricewatcher-worker-darwin-amd64
	@printf " %s\n" "Done!"
	@printf "%s" "Building linux/amd64..."
	@env GOOS=linux GOARCH=amd64 go build -o builds/pricewatcher-worker-linux-amd64
	@printf " %s\n" "Done!"
	@printf "%s" "Building windows/amd64..."
	@env GOOS=windows GOARCH=amd64 go build -o builds/pricewatcher-worker-windows-amd64.exe
	@printf " %s\n" "Done!"

test:
	@echo "Running gotest..."
	@gotest ./... -coverprofile=coverage.out -count=1

clean: imports tidy format
	@printf "%s" "Cleaning project..."
	@rm -rf builds
	@printf " %s\n" "Done!"

check: test
	@golangci-lint run

imports:
	@goimports -local github.com/golangci/golangci-lint -w .

tidy:
	@go mod tidy

format:
	@gofmt -s -w .