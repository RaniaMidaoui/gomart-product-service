FROM golang:1.20-alpine as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

COPY . .
RUN go build -o main ./cmd

FROM builder as test
RUN go test ./... -v

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 50052

CMD ["./main"]
