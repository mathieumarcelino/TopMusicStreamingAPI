FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY .env .

RUN apk --no-cache add \
    ca-certificates \
    tzdata

ENV TZ=Europe/Paris

RUN mkdir -p /app/json /app/data

RUN adduser -D musicapi && \
    chown -R musicapi:musicapi /app

USER musicapi

CMD ["./main"]
