FROM golang:latest

WORKDIR go/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

CMD ["./app"]