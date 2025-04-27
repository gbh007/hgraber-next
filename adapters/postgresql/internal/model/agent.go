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
		"has_hproxy",
		"priority",
		"create_at",
	}
}

func AgentScanner(agent *core.Agent) RowScanner {
	return func(rows pgx.Rows) error {
		var (
			u string
		)

		err := rows.Scan(
			&agent.ID,
			&agent.Name,
			&u,
			&agent.Token,
			&agent.CanParse,
			&agent.CanParseMulti,
			&agent.CanExport,
			&agent.HasFS,
			&agent.HasHProxy,
			&agent.Priority,
			&agent.CreateAt,
		)
		if err != nil {
			return fmt.Errorf("scan to model: %w", err)
		}

		addr, err := url.Parse(u)
		if err != nil {
			return fmt.Errorf("parse addr: %w", err)
		}

		agent.Addr = *addr

		return nil
	}
}
