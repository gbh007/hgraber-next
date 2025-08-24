package label

import "github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"

type LabelRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *LabelRepo {
	return &LabelRepo{
		Repository: repo,
	}
}
