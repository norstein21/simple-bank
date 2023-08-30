postgres:

migrateup:
	migrate -path db/migration -database "postgresql://postgres:habie@localhost:5430/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:habie@localhost:5430/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres migrateup migratedown sqlc test