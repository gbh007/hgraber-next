package config

import "time"

type Application struct {
	Debug              ApplicationDebug `yaml:"debug" envconfig:"DEBUG"`
	MetricScrapePeriod time.Duration    `yaml:"metric_scrape_period" envconfig:"METRIC_SCRAPE_PERIOD"`
	ServiceName        string           `yaml:"service_name" envconfig:"SERVICE_NAME"`
	TraceEndpoint      string           `yaml:"trace_endpoint" envconfig:"TRACE_ENDPOINT"`
}

func ApplicationDefault() Application {
	return Application{
		Debug:              ApplicationDebugDefault(),
		MetricScrapePeriod: 10 * time.Second,
		ServiceName:        "hgraber-next",
	}
}

type ApplicationDebug struct {
	Logs      bool `yaml:"logs" envconfig:"LOGS"`
	DB        bool `yaml:"db" envconfig:"DB"`
	APIServer bool `yaml:"api_server" envconfig:"API_SERVER"`
	APIAgent  bool `yaml:"api_agent" envconfig:"API_AGENT"`
}

func ApplicationDebugDefault() ApplicationDebug {
	return ApplicationDebug{
		Logs:      false,
		DB:        false,
		APIServer: false,
		APIAgent:  false,
	}
}
