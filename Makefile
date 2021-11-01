build:
	go build

migrate-latest:
	migrate -source file://migrations -database sqlite3://test.db up

drop:
	rm test.db

migrate-down:
	migrate -source file://migrations -database sqlite3://test.db down 1
