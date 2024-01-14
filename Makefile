network:
	docker network create bank-network
postgres:
	docker run --name postgres --network bank-network -p 5435:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5435/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go jomabank/db/sqlc Store

.PHONY: network postgres createdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock