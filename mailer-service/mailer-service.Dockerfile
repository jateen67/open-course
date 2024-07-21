FROM golang:1.20.4-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o mailerExec ./cmd/api

RUN chmod +x /app/mailerExec

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/mailerExec /app
COPY --from=builder /app/templates /templates

CMD [ "/app/mailerExec" ]