FROM golang:1.21.1 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o transact-ease ./cmd/api/main.go

FROM alpine:3

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/transact-ease /transact-ease

CMD ["/transact-ease"]
