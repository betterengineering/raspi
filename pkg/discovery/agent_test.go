// Copywright © 2018 Mark Spicer

package discovery_test

import (
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/mdns"
	"github.com/lodge93/raspi/pkg/discovery"
	"github.com/stretchr/testify/suite"
)

type AgentTestSuite struct {
	suite.Suite
	agent *discovery.Agent
}

func TestAgentTestSuite(t *testing.T) {
	suite.Run(t, new(AgentTestSuite))
}

func (suite *AgentTestSuite) SetupTest() {
	agent, err := discovery.NewAgent()
	if err != nil {
		suite.FailNow("could not setup test suite:", err)
	}

	suite.agent = agent
}

func (suite *AgentTestSuite) TeardownTest() {
	err := suite.agent.Shutdown()
	suite.Nil(err)
}

func (suite *AgentTestSuite) TestAgentDiscovery() {
	var wg sync.WaitGroup
	agents := make(chan *mdns.ServiceEntry, 4)
	found := false

	host, err := os.Hostname()
	if err != nil {
		suite.FailNow("could not setup test suite:", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for agent := range agents {
			if strings.TrimSuffix(agent.Host, ".") == host {
				found = true
			}
		}
	}()

	mdns.Lookup(discovery.AgentServiceName, agents)

	close(agents)
	wg.Wait()
	suite.True(found)
}
