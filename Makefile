run:
	go run main.go

build:
	go build -o main main.go
	./main

clean:
	rm -f main

build_all_os: build_linux build_windows build_darwin

build_linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -o main_linux main.go

build_windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o main_windows.exe main.go

build_darwin:
	@echo "Building for macOS (Darwin)..."
	GOOS=darwin GOARCH=amd64 go build -o main_darwin main.go

clean_all:
	rm -f main_linux main_windows.exe main_darwin
