package urlmirror

import "github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"

type URLMirrorRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *URLMirrorRepo {
	return &URLMirrorRepo{
		Repository: repo,
	}
}
