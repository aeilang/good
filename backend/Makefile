new_migrate:
	@migrate create -ext sql -dir db/migration ${name}

migration_up: migration_clean
	@migrate -database="postgres://lang:password@localhost:5432/test_db?sslmode=disable" \
	-path="db/migration" up

migration_down:
	@migrate -database="postgres://lang:password@localhost:5432/test_db?sslmode=disable" \
	-path="db/migration" down

migration_clean:
	@migrate -database="postgres://lang:password@localhost:5432/test_db?sslmode=disable" \
	-path="db/migration" drop -f

init_postgres:
	@docker run -d \
	--name test-postgres \
	-p 5432:5432 \
	-e POSTGRES_USER=lang \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_DB=test_db \
	-v /custom/mount:/var/lib/postgresql/data \
	postgres:alpine

sqlc:
	@sqlc generate