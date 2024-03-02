set dotenv-load

default:
  just --list

css:
	bun tailwindcss -i ./web/src/style.css -o ./web/dist/style.css --watch

ts:
	bun build web/src/main.ts --outdir web/dist --watch

templ:
	templ generate --watch

start:
    wgo run cmd/main.go

format-ts:
	prettier --write '**/*.{js,jsx,ts,tsx,css,html,yml,yaml,json,md}'

format-go:
	gofmt -w -l .

format:
	just format-ts
	just format-go

lint-ts:
	eslint . --ext .ts,.tsx --fix

lint-go:
	golangci-lint run

lint:
	just lint-ts
	just lint-go

pre:
	pre-commit run --all-files

test:
	gotestsum --format testname

build:
	go build -o $BINARY_PATH ./cmd && chmod +x $BINARY_PATH

docker-build:
	docker build --platform linux/amd64 -t $DOCKER_IMAGE_NAME .

docker-run: docker-build
	docker run -p $PORT:$PORT $DOCKER_IMAGE_NAME

docker-push: docker-build
	docker push $DOCKER_IMAGE_NAME
