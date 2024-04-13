FROM golang:1.22

WORKDIR /app

COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /main ./cmd/main.go

CMD ["/main"]