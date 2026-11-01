package deadhash

import "github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/repository"

type DeadHashRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *DeadHashRepo {
	return &DeadHashRepo{
		Repository: repo,
	}
}
