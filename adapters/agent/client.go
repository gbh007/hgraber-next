package agent

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/adapters/agent/internal/adapter"
	"github.com/gbh007/hgraber-next/domain/core"
)

type Client struct {
	agents       map[uuid.UUID]*adapter.Adapter
	agentMutex   *sync.RWMutex
	agentTimeout time.Duration
}

func New(agents []core.Agent, agentTimeout time.Duration) (*Client, error) {
	client := &Client{
		agents:       make(map[uuid.UUID]*adapter.Adapter, len(agents)),
		agentMutex:   &sync.RWMutex{},
		agentTimeout: agentTimeout,
	}

	for _, agent := range agents {
		err := client.SetAgent(agent)
		if err != nil {
			return nil, fmt.Errorf("agent %s: %w", agent.ID.String(), err)
		}
	}

	return client, nil
}

func (c *Client) SetAgent(agent core.Agent) error {
	c.agentMutex.Lock()
	defer c.agentMutex.Unlock()

	a, err := adapter.New(agent.Addr.String(), agent.Token, c.agentTimeout)
	if err != nil {
		return err
	}

	// TODO: проверить отсутствие утечек соединений
	c.agents[agent.ID] = a

	return nil
}

func (c *Client) DeleteAgent(id uuid.UUID) error {
	c.agentMutex.Lock()
	defer c.agentMutex.Unlock()

	_, ok := c.agents[id]
	if !ok {
		return core.AgentNotFoundError
	}

	delete(c.agents, id)

	return nil
}

func (c *Client) getAdapter(id uuid.UUID) (*adapter.Adapter, error) {
	c.agentMutex.RLock()
	defer c.agentMutex.RUnlock()

	a, ok := c.agents[id]
	if !ok || a == nil {
		return nil, core.AgentNotFoundError
	}

	return a, nil
}
