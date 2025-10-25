package model

import (
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5"

	"github.com/gbh007/hgraber-next/domain/core"
)

var AgentTable Agent

type Agent struct{}

func (Agent) Name() string {
	return "agents"
}

func (Agent) ColumnID() string            { return "id" }
func (Agent) ColumnName() string          { return "name" }
func (Agent) ColumnAddr() string          { return "addr" }
func (Agent) ColumnToken() string         { return "token" }
func (Agent) ColumnCanParse() string      { return "can_parse" }
func (Agent) ColumnCanParseMulti() string { return "can_parse_multi" }
func (Agent) ColumnCanExport() string     { return "can_export" }
func (Agent) ColumnHasFS() string         { return "has_fs" }
func (Agent) ColumnHasHProxy() string     { return "has_hproxy" }
func (Agent) ColumnPriority() string      { return "priority" }
func (Agent) ColumnCreateAt() string      { return "create_at" }

func (a Agent) Columns() []string {
	return []string{
		a.ColumnID(),
		a.ColumnName(),
		a.ColumnAddr(),
		a.ColumnToken(),
		a.ColumnCanParse(),
		a.ColumnCanParseMulti(),
		a.ColumnCanExport(),
		a.ColumnHasFS(),
		a.ColumnHasHProxy(),
		a.ColumnPriority(),
		a.ColumnCreateAt(),
	}
}

func (Agent) Scanner(agent *core.Agent) RowScanner {
	return func(rows pgx.Rows) error {
		var u string

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
