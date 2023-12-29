# Include environment variables from .env file
include .env
export

start:
	@echo "Starting the app..."
	go run main.go

watch:
	@echo "Starting the app in watch mode..."
	air

build:
	@echo "Building the app..."
	go build -o $(BINARY_PATH) && chmod +x $(BINARY_PATH)

format:
	@echo "Running go fmt..."
	go fmt ./...

test:
	@echo "Running tests..."
	gotestsum --format testname

prettier:
	@echo "Running prettier..."
	prettier --write .

pre:
	@echo "Running pre-commit..."
	pre-commit run --all-files


docker-build:
	@echo "Building Docker image..."
	docker build --platform linux/amd64 -t $(DOCKER_IMAGE_NAME) .

docker-run: docker-build
	@echo "Running Docker image..."
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE_NAME)

# Push the Docker image
docker-push: docker-build
	@echo "Pushing Docker image to Docker Hub..."
	docker push $(DOCKER_IMAGE_NAME)
