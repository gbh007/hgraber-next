package page

import "github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"

type PageRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *PageRepo {
	return &PageRepo{
		Repository: repo,
	}
}
