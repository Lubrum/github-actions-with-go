FROM golang:1.26.3-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/api .

FROM alpine:3.22 AS runtime

RUN apk add --no-cache ca-certificates tzdata \
	&& addgroup -S app \
	&& adduser -S -G app app

WORKDIR /app

COPY --from=build /out/api ./api
COPY --chown=app:app templates ./templates
COPY --chown=app:app assets ./assets

ENV HOST=postgres \
	DBPORT=5432 \
	USER=root \
	DBNAME=root \
	GIN_MODE=release

EXPOSE 8080

USER app

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
	CMD wget -q --spider http://127.0.0.1:8080/health || exit 1

CMD ["./api"]
