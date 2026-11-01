package srvconfig

import "time"

const (
	DefaultParsingBookTimeout  = time.Minute * 5
	DefaultParsingAgentTimeout = time.Minute * 10

	DefaultMEtricScrapePeriodMainInfo      = 10 * time.Second
	DefaultMEtricScrapePeriodBookStatistic = time.Hour
)
