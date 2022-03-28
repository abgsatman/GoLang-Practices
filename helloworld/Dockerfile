FROM golang:1.18

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-go

EXPOSE 8080

CMD [ "/docker-go" ]