-- migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank" -verbose down

DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transfers;
DROP TABLE IF EXISTS accounts;