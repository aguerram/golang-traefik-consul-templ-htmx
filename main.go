package main

import (
	"context"
	"fmt"
	"github.com/aguerram/gtcth/config"
	"github.com/aguerram/gtcth/internal"
	"github.com/aguerram/gtcth/internal/api"
	"github.com/aguerram/gtcth/internal/web"
	log "github.com/sirupsen/logrus"
	"time"
)

var gracefulShutdowns []internal.GracefulShutdownHandler

func main() {
	gracefulShutdowns = make([]internal.GracefulShutdownHandler, 0, 3)
	config.InitLogger()

	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Application panicked: %v", r)
			handlePanic()
		}
	}()

	env := config.InitializeAppEnv()

	connection, closeConnection, err := config.NewDatabaseConnection(env)
	if err != nil {
		log.Error("Error connecting to database")
		log.Fatal(err)
	}
	gracefulShutdowns = append(gracefulShutdowns, closeConnection)

	server, shutdownHttpServer := internal.StartHttpServer(env)
	gracefulShutdowns = append(gracefulShutdowns, shutdownHttpServer)

	//initialize api
	api.InitializeApi(env, connection, server.Group("/api/v1"))

	//initialize web app
	web.InitializeWebApp(env, connection, server.Group("/"))

	//register with consuls
	deregisterConsul := config.RegisterServiceWithConsul(env)
	gracefulShutdowns = append(gracefulShutdowns, deregisterConsul)

	log.Infof("Server started on port %d", env.Port)

	performGracefulShutdown()

	//start http server
	server.Listen(fmt.Sprintf(":%d", env.Port))

}

func handlePanic() {
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()
	for _, handler := range gracefulShutdowns {
		if handler == nil {
			continue
		}
		handler(ctx)
	}
}

func performGracefulShutdown() {
	log.Info("Performing graceful shutdown")
	internal.HandleGracefulShutdowns(gracefulShutdowns...)
}
