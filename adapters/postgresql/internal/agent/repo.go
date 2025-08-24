package agent

import "github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"

type AgentRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *AgentRepo {
	return &AgentRepo{
		Repository: repo,
	}
}
