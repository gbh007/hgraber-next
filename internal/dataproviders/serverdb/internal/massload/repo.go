package massload

import "github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/repository"

type MassloadRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *MassloadRepo {
	return &MassloadRepo{
		Repository: repo,
	}
}
