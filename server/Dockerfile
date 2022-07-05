FROM golang:1.18.2 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex
EXPOSE 8080
CMD reflex -g '*.go' go run server.go --start-service