docker-up:
	docker-compose -f infra/local/docker-compose.yaml up

migrate-up:
	go run cmd/migration/main.go -up

migrate-down:
	go run cmd/migration/main.go -down

seed:
	go run cmd/seed/main.go

run-tests:
	ENVIRONMENT=test go test ./...