package config

import "time"

type Application struct {
	Debug              bool          `yaml:"debug" envconfig:"DEBUG"`
	MetricScrapePeriod time.Duration `yaml:"metric_scrape_period" envconfig:"METRIC_SCRAPE_PERIOD"`
	ServiceName        string        `yaml:"service_name" envconfig:"SERVICE_NAME"`
	TraceEndpoint      string        `yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
}

func ApplicationDefault() Application {
	return Application{
		Debug:              false,
		MetricScrapePeriod: 10 * time.Second,
		ServiceName:        "hgraber-next",
	}
}
