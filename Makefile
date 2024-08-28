DSN := $(shell godotenv -f .env -- sh -c 'echo $$DSN')


run:
	air

consul-dev:
	consul agent -dev

traefik-dev:
	traefik --configFile=assets/traefik/traefik.yml

generate-certificate:
	mkcert -cert-file assets/ssl/cert.crt -key-file assets/ssl/key.key my-app.local

templ:
	templ generate

db-generate:
	sqlc generate -f ./internal/db/base/sqlc.yaml

db-migrate-up:
	@godotenv -f .env -- migrate -path ./internal/db/base/migrations -database $(DSN) up

db-create-migration:
	migrate create -ext sql -dir internal/db/base/migrations -seq $(name)
ifndef name
	$(error name is not set. Usage: make db-create-migration name=<migration_name>)
endif