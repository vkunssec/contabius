FROM hugomods/hugo:latest as docs
WORKDIR /docs
ARG HUGO_BASEURL="https://contabius-h7jpgeybda-uc.a.run.app/docs"
ENV HUGO_BASEURL=${HUGO_BASEURL}
COPY /docs /docs
RUN hugo --gc --minify --forceSyncStatic

FROM golang:alpine AS builder
WORKDIR /app
COPY . /app/
RUN apk add ca-certificates dumb-init
RUN GOPROXY="https://goproxy.io,direct" go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o application ./cmd/main.go

FROM gcr.io/distroless/static AS compiler
WORKDIR /app
COPY --from=builder /app/.en[v] .
COPY --from=builder /app/application .
COPY --from=builder /usr/bin /usr/bin
COPY --from=docs /docs/public /app/docs
ENTRYPOINT ["/usr/bin/dumb-init", "--"]
EXPOSE 8080
CMD ["./application"]
