TOKEN=

DEV_PORT=:5000
PROD_PORT=:8080

MOCK_SOURCE=./internal/services/service.go
MOCK_DESTINATION=./internal/services/mock/mock.go

run:
	@docker compose up db -d
	@POSTGRES_HOST=localhost go run cmd/app/main.go

# DEV COMMANDS
signup:
	@curl -v -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"name":"larry", "username":"larr", "password":"123"}' http://localhost${DEV_PORT}/auth/sign-up

signin:
	@curl -v -c cookie.txt -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"username":"larr", "password":"123"}' http://localhost${DEV_PORT}/auth/sign-in

refresh:
	@curl -v -b ./cookie.txt -c ./cookie.txt -X POST http://localhost${DEV_PORT}/auth/refresh-token

api:
	@curl -v -b ./cookie.txt -X GET http://localhost${DEV_PORT}/api

logout:
	@curl -v -b ./cookie.txt -X POST http://localhost${DEV_PORT}/auth/logout
	@rm cookie.txt

# PROD COMMANDS
signupProd:
	@curl -v -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"name":"larry", "username":"larr", "password":"123"}' http://localhost${PROD_PORT}/auth/sign-up

signinProd:
	@curl -v -c cookie.txt -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"username":"larr", "password":"123"}' http://localhost${PROD_PORT}/auth/sign-in

refreshProd:
	@curl -v -b ./cookie.txt -c ./cookie.txt -X POST http://localhost${PROD_PORT}/auth/refresh-token

apiProd:
	@curl -v -b ./cookie.txt -X GET http://localhost${PROD_PORT}/api

logoutProd:
	@curl -v -b ./cookie.txt -X POST http://localhost${PROD_PORT}/auth/logout
	@rm cookie.txt


# TESTS
coverage:
	@docker compose up db -d
	@POSTGRES_HOST=localhost go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@rm coverage.out
	@docker compose down


# MOCKGEN
mockgen:
	@mockgen -source=${MOCK_SOURCE} -destination=${MOCK_DESTINATION}

# MIGRATIONS
up:
	@goose -dir ./internal/repository/migrations postgres "user=user password=password dbname=db sslmode=disable" up

down:
	@goose -dir ./internal/repository/migrations  postgres "user=user password=password dbname=db sslmode=disable" down