# Build stage
FROM golang:1.23.1 AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# binary build
RUN go build -trimpath -ldflags "-w -s" -o app

# Deployment stage
FROM debian:bullseye-slim as deploy

COPY --from=deploy-bullder /app/app .

CMD ["./app"]

# Development stage with Air for live reloading

FROM golang:1.23.1 as dev
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]