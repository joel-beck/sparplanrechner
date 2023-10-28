APP_NAME=sparplanrechner

start:
	go run main.go

test:
	gotestsum --format testname

build:
	go build -o bin/$(APP_NAME) main.go

