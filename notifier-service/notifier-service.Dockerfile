FROM golang:1.20.4-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o notifierExec ./cmd/api

RUN chmod +x /app/notifierExec

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/notifierExec /app

CMD [ "/app/notifierExec" ]