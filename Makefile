build:
	go build

migrate-latest:
	migrate -source file://migrations -database $$(cat db.json | jq -r '.database') up

drop:
	rm test.db

migrate-down:
	migrate -source file://migrations -database $$(cat db.json | jq -r '.database') down 1
