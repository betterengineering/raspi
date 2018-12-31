// Copywright © 2018 Mark Spicer

package main

import (
	"log"

	"github.com/lodge93/raspi/pkg/discovery"
)

func main() {
	agent, err := discovery.NewAgent()
	if err != nil {
		log.Fatal(err)
	}
	defer agent.Shutdown()

	log.Println("listening")

	select {}
}
