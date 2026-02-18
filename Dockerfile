# --------- Stage 1: Build ---------
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 👇 build apontando para ./cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o portal ./cmd

# --------- Stage 2: Runtime ---------
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/portal .

EXPOSE 8080

CMD ["./portal"]
