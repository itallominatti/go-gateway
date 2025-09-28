go-migrate:
	migrate -database "postgres://postgres:postgres@localhost:5432/gateway?sslmode=disable" -path migrations up

go-run:
	go run cmd/app/main.go