// Copywright © 2018 Mark Spicer

package main

import (
	"log"
	"sync"
	"time"

	"github.com/hashicorp/mdns"
	"github.com/lodge93/raspi/pkg/discovery"
)

func main() {
	var wg sync.WaitGroup

	agents := make(chan *mdns.ServiceEntry, 100)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for agent := range agents {
			log.Println(agent)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			mdns.Lookup(discovery.AgentServiceName, agents)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
}
