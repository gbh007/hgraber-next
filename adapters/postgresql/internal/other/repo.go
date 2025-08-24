package other

import "github.com/gbh007/hgraber-next/adapters/postgresql/internal/repository"

// OtherRepo - представляет из себя сборник того чему не нашлось места в других модулях.
type OtherRepo struct {
	*repository.Repository
}

func New(repo *repository.Repository) *OtherRepo {
	return &OtherRepo{
		Repository: repo,
	}
}
