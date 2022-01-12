FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-person -ldflags="-w -s" ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/go-person /usr/bin/
EXPOSE 8000
ENTRYPOINT ["go-person"]