DB_URL=postgres://root:UkL7kUIEgz30PEEpQIOLlvkr2eJiVElT@dpg-ceitaisgqg4dlfd4uk1g-a/webbanhang_70l2

network:
	docker network create bank-network

postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb -U root --owner=root webbanhang

dropdb:
	docker exec -it postgres dropdb webbanhang

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server
	