package config

import "time"

type Metric struct {
	ScrapePeriod MetricScrapePeriod `toml:"scrape_period" yaml:"scrape_period" envconfig:"SCRAPE_PERIOD"`
}

type MetricScrapePeriod struct {
	MainInfo      time.Duration `toml:"main_info" yaml:"main_info" envconfig:"MAIN_INFO"`
	BookStatistic time.Duration `toml:"book_statistic" yaml:"book_statistic" envconfig:"BOOK_STATISTIC"`
}

func MetricDefault() Metric {
	return Metric{
		ScrapePeriod: MetricScrapePeriod{
			MainInfo:      DefaultMEtricScrapePeriodMainInfo,
			BookStatistic: DefaultMEtricScrapePeriodBookStatistic,
		},
	}
}

func (m Metric) Enabled() bool {
	return m.ScrapePeriod.MainInfo > 0 || m.ScrapePeriod.BookStatistic > 0
}

func (m Metric) MainInfo() time.Duration {
	return m.ScrapePeriod.MainInfo
}

func (m Metric) BookStatistic() time.Duration {
	return m.ScrapePeriod.BookStatistic
}
