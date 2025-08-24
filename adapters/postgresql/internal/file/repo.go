package file

import "github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"

type FileRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *FileRepo {
	return &FileRepo{
		Repository: repo,
	}
}
