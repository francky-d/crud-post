FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod  ./
RUN go mod download

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]