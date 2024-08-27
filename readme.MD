# Step 1 start consul

```
consul agent -dev
```

> Consul UI available at http://localhost:8500

# Step 2 start traefik

```
traefik --configFile=assets/traefik/traefik.yml
```

### To generate ssl certificate use
```
mkcert -install
mkcert my-app.local
```

### Project structure
<pre>
/myapp
├── /cmd
│   └── /myapp
│       └── main.go
├── /config
│   ├── config.go
│   └── env.go
├── /internal
│   ├── /api
│   │   ├── /handlers
│   │   │   ├── user_handler.go
│   │   │   └── auth_handler.go
│   │   ├── /services
│   │   │   ├── user_service.go
│   │   │   └── auth_service.go
│   │   ├── /request
│   │   │   └── user_request.go
│   │   ├── /response
│   │   │   └── user_response.go
│   │   ├── /types
│   │   │   └── api_types.go
│   │   └── router.go
│   ├── /web
│   │   ├── /handlers
│   │   │   ├── home_handler.go
│   │   │   └── dashboard_handler.go
│   │   ├── /services
│   │   │   ├── home_service.go
│   │   │   └── dashboard_service.go
│   │   ├── /templates
│   │   │   ├── base.templ
│   │   │   ├── home.templ
│   │   │   └── dashboard.templ
│   │   ├── /static
│   │   │   ├── css
│   │   │   │   └── styles.css
│   │   │   └── js
│   │   │       └── scripts.js
│   │   ├── /request
│   │   │   └── web_request.go
│   │   ├── /response
│   │   │   └── web_response.go
│   │   ├── /types
│   │   │   └── web_types.go
│   │   └── router.go
│   ├── /db
│   │   ├── db.go
│   │   ├── migrations/
│   │   └── models/
│   ├── /middleware
│   │   ├── auth_middleware.go
│   │   └── logging_middleware.go
│   └── /consul
│       └── consul_service.go
├── /pkg
│   ├── /utils
│   │   ├── http_utils.go
│   │   └── string_utils.go
│   └── /thirdparty
│       ├── thirdparty_api.go
│       └── external_service.go
├── /scripts
│   ├── migrate.sh
│   └── start.sh
├── /assets
│   ├── /ssl
│   │   ├── myapp.local.crt
│   │   └── myapp.local.key
│   └── /images
│       └── logo.png
├── go.mod
└── go.sum
</pre>