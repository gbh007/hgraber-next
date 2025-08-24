package book

import "github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"

type BookRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *BookRepo {
	return &BookRepo{
		Repository: repo,
	}
}
