package agent

import "github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/repository"

type AgentRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *AgentRepo {
	return &AgentRepo{
		Repository: repo,
	}
}
