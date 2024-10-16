FROM golang:1.18.2-alpine3.15 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /output/rabbit-publisher

FROM alpine:3.15

COPY --from=builder /output/rabbit-publisher /rabbit-publisher

EXPOSE 8080

CMD ["/rabbit-publisher"]
