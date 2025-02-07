package agentmodel

import (
	"time"

	"github.com/gbh007/hgraber-next/domain/core"
)

type AgentStatus struct {
	StartAt   time.Time
	IsOK      bool
	IsWarning bool
	IsError   bool
	Problems  []AgentStatusProblem
}

type AgentStatusProblem struct {
	IsInfo    bool
	IsWarning bool
	IsError   bool
	Details   string
}

type AgentWithStatus struct {
	Agent       core.Agent
	Status      AgentStatus
	IsOffline   bool
	StatusError string
}
