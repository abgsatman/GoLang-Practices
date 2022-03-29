FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o ./out/go-app .

# Run the binary program produced by `go install`
CMD ["./out/go-app"]