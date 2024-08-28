package main

import (
	"fmt"
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal"
	"github.com/aguerram/gtcth/internal/api"
	"github.com/aguerram/gtcth/internal/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	config.InitLogger()
	env := config.InitializeAppEnv()

	_, closeConnection, err := config.NewDatabaseConnection(env)
	if err != nil {
		log.Error("Error connecting to database")
		log.Fatal(err)
	}

	server, shutdownHttpServer := internal.StartHttpServer(env)

	//initialize api
	api.InitializeApi(env, server.Group("/api/v1"))

	//initialize web app
	web.InitializeWebApp(env, server.Group("/"))

	//register with consul
	deregisterConsul := config.RegisterServiceWithConsul(env)
	internal.HandleGracefulShutdowns(deregisterConsul, shutdownHttpServer, closeConnection)

	log.Infof("Server started on port %d", env.Port)
	//start http server
	server.Listen(fmt.Sprintf(":%d", env.Port))
}
