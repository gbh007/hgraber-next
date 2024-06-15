package model

import (
	"time"

	"github.com/google/uuid"

	"hgnext/internal/entities"
)

type Agent struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Addr      string    `db:"addr"`
	Token     string    `db:"token"`
	CanParse  bool      `db:"can_parse"`
	CanExport bool      `db:"can_export"`
	Priority  int       `db:"priority"`
	CreateAt  time.Time `db:"create_at"`
}

func (a Agent) ToEntity() (entities.Agent, error) {
	id, err := uuid.Parse(a.ID)
	if err != nil {
		return entities.Agent{}, err
	}

	return entities.Agent{
		ID:        id,
		Name:      a.Name,
		Addr:      a.Addr,
		Token:     a.Token,
		CanParse:  a.CanParse,
		CanExport: a.CanExport,
		Priority:  a.Priority,
		CreateAt:  a.CreateAt,
	}, nil
}
