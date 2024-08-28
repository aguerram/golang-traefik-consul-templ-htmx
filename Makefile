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
	sqlc generate -f ./assets/sqlc/sqlc.yaml

db-push-tags:
	# add more tags
	sqlc push -f assets/sqlc/sqlc.yaml

db-verify:
	# add more schema to verify
	sqlc verify -f assets/sqlc/sqlc.yaml
