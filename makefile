TOKEN=

PORT=:5000

run:
	@go run cmd/app/main.go
	
up:
	@goose -dir ./internal/repository/migrations postgres "user=user password=password dbname=db sslmode=disable" up

down:
	@goose -dir ./internal/repository/migrations  postgres "user=user password=password dbname=db sslmode=disable" down

signup:
	@curl -v -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"name":"larry", "username":"larr", "password":"123"}' http://localhost${PORT}/auth/sign-up

signin:
	@curl -v -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"username":"larr", "password":"123"}' http://localhost${PORT}/auth/sign-in

api:
	@curl -v -H 'Accept: application/json' -H 'Authorization: Bearer ${TOKEN}' http://localhost${PORT}/api
