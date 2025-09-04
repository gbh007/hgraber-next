package config

import "time"

type Application struct {
	Metric        Metric    `toml:"metric" yaml:"metric" envconfig:"METRIC"`
	ServiceName   string    `toml:"service_name" yaml:"service_name" envconfig:"SERVICE_NAME"`
	TraceEndpoint string    `toml:"trace_endpoint" yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
	Pyroscope     Pyroscope `toml:"pyroscope" yaml:"pyroscope" envconfig:"PYROSCOPE"`
}

func ApplicationDefault() Application {
	return Application{
		Metric:      MetricDefault(),
		ServiceName: "hgraber-next",
	}
}

type Metric struct {
	ScrapePeriod MetricScrapePeriod `toml:"scrape_period" yaml:"scrape_period" envconfig:"SCRAPE_PERIOD"`
}

func MetricDefault() Metric {
	return Metric{
		ScrapePeriod: MetricScrapePeriod{
			MainInfo:      10 * time.Second,
			BookStatistic: time.Hour,
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

type MetricScrapePeriod struct {
	MainInfo      time.Duration `toml:"main_info" yaml:"main_info" envconfig:"MAIN_INFO"`
	BookStatistic time.Duration `toml:"book_statistic" yaml:"book_statistic" envconfig:"BOOK_STATISTIC"`
}

type Pyroscope struct {
	Endpoint string `toml:"endpoint" yaml:"endpoint" envconfig:"ENDPOINT"`
	Debug    bool   `toml:"debug" yaml:"debug" envconfig:"DEBUG"`
	Rate     int    `toml:"rate" yaml:"rate" envconfig:"RATE"`
}
