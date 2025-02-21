package config

import "time"

type Application struct {
	Debug         ApplicationDebug `yaml:"debug" envconfig:"DEBUG"`
	Metric        Metric           `yaml:"metric" envconfig:"METRIC"`
	ServiceName   string           `yaml:"service_name" envconfig:"SERVICE_NAME"`
	TraceEndpoint string           `yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
	Pyroscope     Pyroscope        `yaml:"pyroscope" envconfig:"PYROSCOPE"`
}

func ApplicationDefault() Application {
	return Application{
		Debug:       ApplicationDebugDefault(),
		Metric:      MetricDefault(),
		ServiceName: "hgraber-next",
	}
}

type ApplicationDebug struct {
	Logs bool `yaml:"logs" envconfig:"LOGS"`
}

func ApplicationDebugDefault() ApplicationDebug {
	return ApplicationDebug{
		Logs: false,
	}
}

type Metric struct {
	ScrapePeriod MetricScrapePeriod `yaml:"scrape_period" envconfig:"SCRAPE_PERIOD"`
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
	MainInfo      time.Duration `yaml:"main_info" envconfig:"MAIN_INFO"`
	BookStatistic time.Duration `yaml:"book_statistic" envconfig:"BOOK_STATISTIC"`
}

type Pyroscope struct {
	Endpoint string `yaml:"endpoint" envconfig:"ENDPOINT"`
	Debug    bool   `yaml:"debug" envconfig:"DEBUG"`
	Rate     int    `yaml:"rate" envconfig:"RATE"`
}
