postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret --restart=always -d postgres:15.2

createdb:
	docker exec -it postgres createdb --username=root --owner=root parking

dropdb:
	docker exec -it postgres dropdb parking

run:
	go run cmd/main.go

lint:
	golangci-lint run --timeout 10m ./...

test:
	go test -v -coverprofile=cover.out ./...
	go tool cover -func=cover.out

build:
	docker build --platform=linux/amd64 -t sugamdocker35/lms-api:v2.0.6 .

push:
	docker push sugamdocker35/lms-api:v2.0.6

swag:
	swag init -g cmd/main.go -o ./docs --ot go --parseInternal true


proto-gen:
	protoc --go_out=. --go-grpc_out=. proto/notification/notification.proto

reload:
	air

extension:
	CREATE EXTENSION IF NOT EXISTS pg_trgm;

up:
	docker-compose up -d

down:
	docker-compose down

#Check Docker Desktop disk usage
chec-usage:
	docker system df 

#Delete build cache (safe, biggest gain):
build-cache:
	docker builder prune -af
