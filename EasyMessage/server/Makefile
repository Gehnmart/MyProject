.PHONY: clean all

PORT = 4000
ENVIROMENT = development


BODY="'{"name": "egor", "email":"egor.com", "password":"112"}'"
all: server

server:
	go run ./cmd/api --port ${PORT} --env ${ENVIROMENT}

test_endpoints:
	curl -i localhost:${PORT}/v1/healthcheck/
	curl -i -X POST localhost:${PORT}/v1/movies
	curl -i -X GET localhost:${PORT}/v1/movies/111

	curl -i -d "$(BODY)" localhost:${PORT}/v1/user
	
migrate_up:
	migrate -path ./migrations -database=$EASYMESSAGE_DB_DSN -verbose up