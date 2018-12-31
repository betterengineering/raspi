// Copywright © 2018 Mark Spicer

// Package discovery provides a client and server for discovering raspberry pis
// via mDNS.
package discovery

import (
	"os"

	"github.com/hashicorp/mdns"
)

const (
	AgentServiceName = "_discoveryd._tcp"
	AgentServiceInfo = "raspberry pi discovery service"
	AgentPort        = 8000
)

type Agent struct {
	mdnsServer *mdns.Server
}

func NewAgent() (*Agent, error) {
	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	info := []string{AgentServiceInfo}
	service, err := mdns.NewMDNSService(host, AgentServiceName, "", "", AgentPort, nil, info)
	if err != nil {
		return nil, err
	}

	server, err := mdns.NewServer(&mdns.Config{Zone: service})
	if err != nil {
		return nil, err
	}

	return &Agent{
		mdnsServer: server,
	}, nil
}

func (agent *Agent) Shutdown() error {
	return agent.mdnsServer.Shutdown()
}
