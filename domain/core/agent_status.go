package core

import "time"

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
	Agent       Agent
	Status      AgentStatus
	IsOffline   bool
	StatusError string
}
