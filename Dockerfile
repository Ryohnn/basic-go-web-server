FROM golang:1.26-alpine3.24 AS dev

WORKDIR /app

RUN apk add --no-cache git
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air"]

FROM golang:1.26-alpine3.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

FROM alpine:3.24 AS prod
WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata && \
    addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/main .

USER appuser:appgroup

EXPOSE 8080
CMD ["./main"]