include .env
export

DEV_PORT=:5000
DOCKER_PORT=:8080

MOCK_SOURCE=./internal/services/services.go
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

# DOCKER COMMANDS
signupDocker:
	@curl -v -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"name":"larry", "username":"larr", "password":"123"}' http://localhost${DOCKER_PORT}/auth/sign-up

signinDocker:
	@curl -v -c cookie.txt -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"username":"larr", "password":"123"}' http://localhost${DOCKER_PORT}/auth/sign-in

refreshDocker:
	@curl -v -b ./cookie.txt -c ./cookie.txt -X POST http://localhost${DOCKER_PORT}/auth/refresh-token

apiDocker:
	@curl -v -b ./cookie.txt -X GET http://localhost${DOCKER_PORT}/api

logoutDocker:
	@curl -v -b ./cookie.txt -X POST http://localhost${DOCKER_PORT}/auth/logout
	@rm cookie.txt


# TESTS
coverage:
	@docker compose up db -d
	@POSTGRES_HOST=localhost go test -coverprofile=coverage.out -coverpkg=./... ./...
	@go tool cover -html=coverage.out
	@rm coverage.out
	@docker compose down


# MOCKGEN
mockgen:
	@mockgen -source=${MOCK_SOURCE} -destination=${MOCK_DESTINATION}

# MIGRATIONS
up:
	@goose -dir ./internal/repository/migrations postgres "host=localhost port=5430 user=user password=password dbname=db sslmode=disable" up

down:
	@goose -dir ./internal/repository/migrations  postgres "host=localhost user=user password=password dbname=db sslmode=disable" down