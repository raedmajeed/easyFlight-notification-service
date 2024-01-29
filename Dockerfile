FROM golang:latest

WORKDIR go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR ./

RUN go build -o app

CMD ["./app"]