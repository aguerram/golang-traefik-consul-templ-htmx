package main

import (
	"fmt"
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal"
	"github.com/aguerram/gtcth/internal/route"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.InitLogger()

	env := config.InitializeAppEnv()

	server, shutdownHttpServer := internal.StartHttpServer(env)
	//register routes
	route.InitializeRoutes(env, server)

	//register with consul
	deregisterConsul := config.RegisterServiceWithConsul(env)
	internal.HandleGracefulShutdowns(deregisterConsul, shutdownHttpServer)

	log.Infof("Server started on port %d", env.Port)
	//start http server
	log.Fatal(server.Listen(fmt.Sprintf(":%d", env.Port)))
}
