# Stage 1: Build the Go application
FROM golang:1.23.1-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o main

# Stage 2: Create a minimal runtime image
FROM ubuntu:23.10
EXPOSE 8000
WORKDIR /app
COPY --from=build /app/main .
COPY ./templates/ templates/
COPY ./assets/ assets/
ENV HOST=postgres \
    DBPORT=5432 \
    USER=root \
    PASSWORD=root \
    DBNAME=root \
    GIN_MODE=release
CMD ["./main"]
