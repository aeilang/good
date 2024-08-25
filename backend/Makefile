new_migrate:
	@migrate create -seq -ext sql -dir db/migrations ${name}

migration_up: migration_clean
	@migrate -database="postgres://lang:password@localhost:5432/test_db?sslmode=disable" \
	-path="db/migrations" up

migration_down:
	@migrate -database="postgres://lang:password@localhost:5432/test_db?sslmode=disable" \
	-path="db/migrations" down

migration_clean:
	@migrate -database="postgres://lang:password@localhost:5432/test_db?sslmode=disable" \
	-path="db/migrations" drop -f

init_postgres:
	@sudo docker run -d \
	--name test-postgres \
	-p 5432:5432 \
	-e POSTGRES_USER=lang \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_DB=test_db \
	-v /custom/mount:/var/lib/postgresql/data \
	postgres:alpine

sqlc:
	@sqlc generate

init_elas:
	docker run -d --name elasticsearch \
	-p 9200:9200 \
	-p 9300:9300 \
	-e "discovery.type=single-node" \
	docker.elastic.co/elasticsearch/elasticsearch:8.5.1

elas2:
	docker run -p 127.0.0.1:9200:9200 -d --name elasticsearch --network elastic-net \
  	-e ELASTIC_PASSWORD=password \
	-e "discovery.type=single-node" \
  	-e "xpack.security.http.ssl.enabled=false" \
  	-e "xpack.license.self_generated.type=trial" \
 	 docker.elastic.co/elasticsearch/elasticsearch:8.15.0

milisearch:
	docker run -d --name milisearch \
	-p 7700:7700 \
	-e MEILI_ENV='development' \
	-e MEILI_MASTER_KEY='masterKey' \
	-v $(pwd)/meili_data:/meili_data \
	getmeili/meilisearch:v1.9

