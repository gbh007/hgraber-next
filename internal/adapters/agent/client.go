package agent

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	"hgnext/internal/adapters/agent/internal/adapter"
	"hgnext/internal/entities"
)

type Client struct {
	agents     map[uuid.UUID]*adapter.Adapter
	agentMutex *sync.RWMutex
}

func New(agents []entities.Agent) (*Client, error) {
	client := &Client{
		agents:     make(map[uuid.UUID]*adapter.Adapter, len(agents)),
		agentMutex: &sync.RWMutex{},
	}

	for _, agent := range agents {
		err := client.SetAgent(agent)
		if err != nil {
			return nil, fmt.Errorf("agent %s: %w", agent.ID.String(), err)
		}
	}

	return client, nil
}

func (c *Client) SetAgent(agent entities.Agent) error {
	c.agentMutex.Lock()
	defer c.agentMutex.Unlock()

	a, err := adapter.New(agent.Addr, agent.Token)
	if err != nil {
		return err
	}

	c.agents[agent.ID] = a

	return nil
}

func (c *Client) DeleteAgent(id uuid.UUID) error {
	c.agentMutex.Lock()
	defer c.agentMutex.Unlock()

	_, ok := c.agents[id]
	if !ok {
		return entities.AgentNotFoundError
	}

	delete(c.agents, id)

	return nil
}

func (c *Client) getAdapter(id uuid.UUID) (*adapter.Adapter, error) {
	c.agentMutex.RLock()
	defer c.agentMutex.RUnlock()

	a, ok := c.agents[id]
	if !ok || a == nil {
		return nil, entities.AgentNotFoundError
	}

	return a, nil
}
