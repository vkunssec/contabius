FROM golang:alpine AS builder
WORKDIR /app
COPY . /app/
RUN apk add dumb-init
RUN apk add --update hugo
RUN GOPROXY="https://goproxy.io,direct" go mod download
RUN cd docs
RUN hugo --gc --minify --forceSyncStatic
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o application ./cmd/main.go

FROM gcr.io/distroless/static AS deploy
WORKDIR /app
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/.en[v] .
COPY --from=builder /app/application .
COPY --from=builder /usr/bin /usr/bin
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
EXPOSE 8080
CMD ["./application"]
