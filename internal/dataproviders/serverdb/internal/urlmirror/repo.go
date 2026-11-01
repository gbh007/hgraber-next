package urlmirror

import "github.com/gbh007/hgraber-next/internal/dataproviders/serverdb/internal/repository"

type URLMirrorRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *URLMirrorRepo {
	return &URLMirrorRepo{
		Repository: repo,
	}
}
