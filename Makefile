APP_NAME := sudoku

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o bin/$(APP_NAME)_darwin_arm64
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)_linux_amd64
	GOOS=linux GOARCH=arm64 go build -o bin/$(APP_NAME)_linux_arm64
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME)_windows_amd64.exe
	GOOS=windows GOARCH=arm64 go build -o bin/$(APP_NAME)_windows_arm64.exe

.PHONY: build
