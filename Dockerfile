FROM --platform=linux/amd64 golang:1.21-alpine

WORKDIR /usr/src/app

EXPOSE ${HOST_PORT}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in
# subsequent builds if they change
COPY go.mod go.sum ./

# download dependencies
RUN go mod download && go mod verify

# copy the rest of the source code
COPY . .

# build the binary
RUN go build -v -o /usr/local/bin/app ./cmd

# run the binary
CMD ["app"]
