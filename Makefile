APP_NAME := sudoku
SRC_DIR := ./src

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)_darwin_amd64 $(SRC_DIR)
	GOOS=darwin GOARCH=arm64 go build -o bin/$(APP_NAME)_darwin_arm64 $(SRC_DIR)
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)_linux_amd64 $(SRC_DIR)
	GOOS=linux GOARCH=arm64 go build -o bin/$(APP_NAME)_linux_arm64 $(SRC_DIR)
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME)_windows_amd64.exe $(SRC_DIR)
	GOOS=windows GOARCH=arm64 go build -o bin/$(APP_NAME)_windows_arm64.exe $(SRC_DIR)

.PHONY: build
