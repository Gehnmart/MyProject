all:
	echo "ti lox"

migrateup:
	migrate -path ./migration -database="$(shell echo $$EASYMESSAGE_DB_DSN)" -verbose up

migratedown:
	migrate -path ./migration -database="$(shell echo $$EASYMESSAGE_DB_DSN)" -verbose down