set dotenv-load

default:
  just --list

start:
    wgo run pkg/cmd/main.go

templ:
	templ generate --watch

css:
	bun tailwindcss -i ./web/src/style.css -o ./web/dist/style.css --watch

ts:
	bun build web/src/main.ts --outdir web/dist --watch

build:
	go build -o $BINARY_PATH && chmod +x $BINARY_PATH

pre:
	pre-commit run --all-files

test:
	gotestsum --format testname

lint-ts:
	eslint . --ext .ts,.tsx --fix

lint-go:
	golangci-lint run

lint:
	lint-ts
	lint-go

format-ts:
	prettier --write '**/*.{js,jsx,ts,tsx,css,html,yml,yaml,json,md}'

format-go:
	gofmt -w -l .

format:
	format-ts
	format-go

docker-build:
	docker build --platform linux/amd64 -t $DOCKER_IMAGE_NAME .

docker-run: docker-build
	docker run -p $PORT:$PORT $DOCKER_IMAGE_NAME

docker-push: docker-build
	docker push $DOCKER_IMAGE_NAME
