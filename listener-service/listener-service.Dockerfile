FROM golang:1.20.4-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/listenerExec ./cmd/api

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/listenerExec .

CMD ["./listenerExec"]