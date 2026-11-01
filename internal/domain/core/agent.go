package core

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Agent struct {
	ID            uuid.UUID
	Name          string
	Addr          url.URL
	Token         string
	CanParse      bool
	CanParseMulti bool
	CanExport     bool
	HasFS         bool
	HasHProxy     bool
	Priority      int
	CreateAt      time.Time
}

type AgentFilter struct {
	CanParse      bool
	CanParseMulti bool
	CanExport     bool
	HasFS         bool
	HasHProxy     bool
}
