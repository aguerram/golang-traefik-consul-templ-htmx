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