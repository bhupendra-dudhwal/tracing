FROM golang:1.23.9-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY . .

RUN go build -o service-c .

# EXPOSE 8003
CMD ["./service-c"]