package model

import (
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

func AgentColumns() []string {
	return []string{
		"id",
		"name",
		"addr",
		"token",
		"can_parse",
		"can_parse_multi",
		"can_export",
		"has_fs",
		"priority",
		"create_at",
	}
}

func AgentScanner(p *core.Agent) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			u string
		)

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&u,
			&p.Token,
			&p.CanParse,
			&p.CanParseMulti,
			&p.CanExport,
			&p.HasFS,
			&p.Priority,
			&p.CreateAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		addr, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("parse addr: %w", err)
		}

		p.Addr = *addr

		return nil
	}
}
