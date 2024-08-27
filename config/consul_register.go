package config

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"os"
)

func RegisterServiceWithConsul(env *AppEnv) func() {
	ipAddress, err := getLocalIP()
	if err != nil {
		log.Fatalf("Failed to get local IP address: %v", err)
	}
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	serviceID := fmt.Sprintf("%s-%d", env.DiscoveryAppName, env.Port)
	serviceName := env.DiscoveryAppName

	tags := []string{
		"traefik.enable=true",
		fmt.Sprintf("traefik.http.routers.%s.rule=Host(`%s`)", env.DiscoveryAppName, env.AppUrl),
		//fmt.Sprintf("traefik.http.routers.%s.entrypoints=websecure", env.DiscoveryAppName),
		fmt.Sprintf("traefik.http.routers.%s.entrypoints=web", env.DiscoveryAppName),
		//fmt.Sprintf("traefik.http.routers.%s.tls.certresolver=default", env.DiscoveryAppName),
		//fmt.Sprintf("traefik.http.services.%s.loadbalancer.server.port=%d", env.DiscoveryAppName, env.Port),
	}
	// > Register service with Consul
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Port:    env.Port,
		Address: ipAddress,
		Check: &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", ipAddress, env.Port),
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "11s",
		},
		Tags: tags,
	}

	consulConfig := api.DefaultConfig()
	consulConfig.Address = env.ConsulAddress
	client, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Failed to register service with Consul: %v", err)
	}

	log.Infof("Service %s registered in Consul", serviceName)

	return func() {
		err := client.Agent().ServiceDeregister(serviceID)
		if err != nil {
			log.Fatalf("Failed to deregister service from Consul: %v", err)
		}
		log.Infof("Service %s deregistered from Consul", serviceName)
	}
}

func getLocalIP() (string, error) {
	//addrs, err := net.InterfaceAddrs()
	//if err != nil {
	//	return "", err
	//}
	//
	//for _, addr := range addrs {
	//	// Check if the address is a valid IP address and not a loopback address
	//	if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
	//		if ipNet.IP.To4() != nil {
	//			return ipNet.IP.String(), nil
	//		}
	//	}
	//}
	//return "", fmt.Errorf("unable to find local IP address")
	return os.Hostname()
}
