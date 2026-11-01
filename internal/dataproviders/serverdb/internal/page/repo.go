package page

import "github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/repository"

type PageRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *PageRepo {
	return &PageRepo{
		Repository: repo,
	}
}
