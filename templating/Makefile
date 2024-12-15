run:
	@echo "Running..."
	go run cmd/main.go

build:
	@echo "Building..."
	go build -ldflags "-s -w" -o bin/main cmd/main.go
	GOOS=windows go build -ldflags "-s -w" -o bin/main.exe cmd/main.go
	GOOS=darwin go build -ldflags "-s -w" -o bin/main-darwin cmd/main.go

clean:
	@echo "Cleaning..."
	rm -rf bin

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run    - Run the application"
	@echo "  build  - Build the application"
	@echo "  clean  - Clean the application"
	@echo "  help   - Display this help message"
	@echo ""