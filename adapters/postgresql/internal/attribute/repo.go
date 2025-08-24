package attribute

import "github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"

type AttributeRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *AttributeRepo {
	return &AttributeRepo{
		Repository: repo,
	}
}
