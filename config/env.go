package config

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type AppEnv struct {
	ConsulAddress    string
	Port             int
	AppName          string
	DiscoveryAppName string
	InstanceUUID     string
	AppUrl           string
}

func InitializeAppEnv(filename ...string) *AppEnv {
	err := godotenv.Load(filename...)
	if err != nil {
		log.Panicf("Failed to load environment variables: %v", err)
	}
	port, err := getAvailablePort()
	if err != nil {
		log.Fatal(err)
	}
	return &AppEnv{
		Port:             port,
		AppName:          os.Getenv("APP_NAME"),
		ConsulAddress:    os.Getenv("CONSUL_ADDRESS"),
		DiscoveryAppName: os.Getenv("DISCOVERY_APP_NAME"),
		InstanceUUID:     uuid.New().String(),
		AppUrl:           os.Getenv("APP_URL"),
	}
}

func getAvailablePort() (int, error) {
	//// Listen on a random port by specifying ":0"
	//listener, err := net.Listen("tcp", ":0")
	//if err != nil {
	//	return 0, err
	//}
	//defer func(listener net.Listener) {
	//	err := listener.Close()
	//	if err != nil {
	//		log.Errorf("Failed to close listener: %v", err)
	//	}
	//}(listener)
	//
	//// Get the port assigned by the system
	//port := listener.Addr().(*net.TCPAddr).Port
	//return port, nil
	getenv := os.Getenv("PORT")
	if getenv == "" {
		return 8080, nil
	}
	return strconv.Atoi(getenv)
}
