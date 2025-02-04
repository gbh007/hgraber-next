package model

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/gbh007/hgraber-next/domain/core"
)

type Agent struct {
	ID            string    `db:"id"`
	Name          string    `db:"name"`
	Addr          string    `db:"addr"`
	Token         string    `db:"token"`
	CanParse      bool      `db:"can_parse"`
	CanParseMulti bool      `db:"can_parse_multi"`
	CanExport     bool      `db:"can_export"`
	HasFS         bool      `db:"has_fs"`
	Priority      int       `db:"priority"`
	CreateAt      time.Time `db:"create_at"`
}

func (a Agent) ToEntity() (core.Agent, error) {
	id, err := uuid.Parse(a.ID)
	if err != nil {
		return core.Agent{}, fmt.Errorf("parse id: %w", err)
	}

	addr, err := url.Parse(a.Addr)
	if err != nil {
		return core.Agent{}, fmt.Errorf("parse addr: %w", err)
	}

	return core.Agent{
		ID:            id,
		Name:          a.Name,
		Addr:          *addr,
		Token:         a.Token,
		CanParse:      a.CanParse,
		CanParseMulti: a.CanParseMulti,
		CanExport:     a.CanExport,
		HasFS:         a.HasFS,
		Priority:      a.Priority,
		CreateAt:      a.CreateAt,
	}, nil
}
