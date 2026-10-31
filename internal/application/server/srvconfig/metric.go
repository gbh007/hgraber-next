package srvconfig

import "time"

type ServerMetric struct {
	ScrapePeriod ServerMetricScrapePeriod `toml:"scrape_period"`
}

type ServerMetricScrapePeriod struct {
	MainInfo      time.Duration `toml:"main_info"`
	BookStatistic time.Duration `toml:"book_statistic"`
}

func ServerMetricDefault() ServerMetric {
	return ServerMetric{
		ScrapePeriod: ServerMetricScrapePeriod{
			MainInfo:      DefaultMEtricScrapePeriodMainInfo,
			BookStatistic: DefaultMEtricScrapePeriodBookStatistic,
		},
	}
}

func (m ServerMetric) Enabled() bool {
	return m.ScrapePeriod.MainInfo > 0 || m.ScrapePeriod.BookStatistic > 0
}

func (m ServerMetric) MainInfo() time.Duration {
	return m.ScrapePeriod.MainInfo
}

func (m ServerMetric) BookStatistic() time.Duration {
	return m.ScrapePeriod.BookStatistic
}
