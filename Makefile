run:
	$ go run cmd/main.go

dev:
	$ air server

build:
	$ CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

test:
	$ go test ./...

docker:
	$ docker compose up --force-recreate --build --remove-orphans

docs-server:
	$ cd docs && hugo server -p 8080

docs-build:
	$ cd docs && hugo --gc --minify --forceSyncStatic
