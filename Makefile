.PHONY: migrate-db
migrate-db:
	migrate -database "postgresql://user:password@127.0.0.1:5432/customers?sslmode=disable" -path db/migrations up

.PHONY: run-server
run-server:
	docker-compose up --build


.PHONY: run-client
run-client:
	go run cmd/client/client.go
