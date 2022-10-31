createdb:
	docker exec -it postgres12 createdb -U root --owner=root webbanhang

dropdb:
	docker exec -it postgres12 dropdb webbanhang

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/webbanhang?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/webbanhang?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/webbanhang?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/webbanhang?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server
	