DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable


postgres:
	docker run --name bank-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it bank-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bank-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database $(DB_URL) -verbose up

migratedown:
	migrate -path db/migration -database $(DB_URL) -verbose down

migrateup1:
	migrate -path db/migration -database $(DB_URL) -verbose up 1

migratedown1:
	migrate -path db/migration -database $(DB_URL) -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/vanhung4499/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration sqlc test server mock