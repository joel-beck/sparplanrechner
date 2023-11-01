# Include environment variables from .env file
include .env
export

start:
	@echo "Starting the app..."
	go run main.go

test:
	@echo "Running tests..."
	gotestsum --format testname

build:
	@echo "Building the app..."
	go build -o $(BINARY_PATH) && chmod +x $(BINARY_PATH)

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

