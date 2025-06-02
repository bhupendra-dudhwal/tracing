FROM golang:1.23.9-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN go build -o service-b .

# EXPOSE 8002
CMD ["./service-b"]