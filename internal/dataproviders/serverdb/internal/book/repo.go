package book

import "github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/repository"

type BookRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *BookRepo {
	return &BookRepo{
		Repository: repo,
	}
}
