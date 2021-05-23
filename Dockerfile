FROM golang:1.16.4 AS builder
COPY go.mod go.sum /app/
WORKDIR /app
RUN sleep 2
RUN go mod download
COPY . /app
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -a -installsuffix cgo -o build ./cmd/main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/build /app
COPY --from=builder /app/migrate.linux-amd64 /app/migrate
COPY .env /app
COPY internal/datastore/postgreSQL/migrations/ /app/migrations/
COPY start.sh /app
EXPOSE 8008 8009 8010
CMD ["/app/build"]
ENTRYPOINT ["/app/start.sh"]
