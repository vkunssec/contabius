run:
	$ go run cmd/main.go

init-dev:
	$ go install github.com/vkunssec/husky@latest
	$ go install github.com/air-verse/air@latest
	$ go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$ go install github.com/swaggo/swag/cmd/swag@latest
	$ make swagger
	$ make dev

dev:
	$ air server

build:
	$ CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

test:
	$ go test ./...

docker:
	$ docker compose up --force-recreate --build --remove-orphans

swagger:
	$ swag init -g cmd/main.go --parseDependency --parseInternal

seed:
	$ go run cmd/seed/main.go

lint:
	$ golangci-lint run ./...
