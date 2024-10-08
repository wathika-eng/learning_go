main:
	go run main.go
build:
	go build -o main main.go
run:
	./main
clean:
	rm main

build_all_os:
	GOOS=linux GOARCH=amd64 go build -o main_linux main.go
	GOOS=windows GOARCH=amd64 go build -o main_windows main.go
	GOOS=darwin GOARCH=amd64 go build -o main_darwin main.go

clean_all:
	rm main_linux main_windows main_darwin