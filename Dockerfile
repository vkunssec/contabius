FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . /app/
RUN apk add ca-certificates dumb-init
RUN GOPROXY="https://goproxy.io,direct" go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go --parseDependency --parseInternal
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o application ./cmd/main.go

FROM gcr.io/distroless/static AS compiler
WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/application .
COPY --from=builder /usr/bin /usr/bin
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
EXPOSE 8080
CMD ["./application"]
