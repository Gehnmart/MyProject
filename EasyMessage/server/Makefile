.PHONY: clean all

PORT = 4000
ENVIROMENT = development


BODY="'{"name": "egor", "email":"eddddd.com", "password":"11adsfasdfas2"}'"
all: server

server:
	go run ./cmd/server --port ${PORT} --env ${ENVIROMENT}

test_endpoints:
#curl -i localhost:${PORT}/v1/healthcheck/
#curl -i -X POST localhost:${PORT}/v1/movies
#curl -i -X GET localhost:${PORT}/v1/movies/111

	curl -i -d "$(BODY)" localhost:${PORT}/signup
	
migrate_up:
	migrate -path ./migrations -database="$(shell echo $$EASYMESSAGE_DB_DSN)" -verbose up
migrate_down:
	migrate -path=./migrations -database="$(shell echo $$EASYMESSAGE_DB_DSN)" -verbose down