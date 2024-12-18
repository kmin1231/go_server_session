# Build stage
FROM golang:1.23.1 AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# binary build
RUN go build -trimpath -ldflags "-w -s" -o app

# Deployment stage
FROM debian:bullseye-slim AS deploy

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# Development stage with Air for live reloading

FROM golang:1.23.1 AS dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]