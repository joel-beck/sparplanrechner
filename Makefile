# Include environment variables from .env file
include .env
export

start:
	go run main.go

watch:
	air

build:
	go build -o $(BINARY_PATH) && chmod +x $(BINARY_PATH)

lint:
	golangci-lint run

format:
	gofmt -w -l .

test:
	gotestsum --format testname


pre:
	pre-commit run --all-files


docker-build:
	docker build --platform linux/amd64 -t $(DOCKER_IMAGE_NAME) .

docker-run: docker-build
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE_NAME)

docker-push: docker-build
	docker push $(DOCKER_IMAGE_NAME)
