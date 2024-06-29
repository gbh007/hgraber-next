package application

import "github.com/kelseyhightower/envconfig"

type Config struct {
	PostgreSQLConnection  string `envconfig:"POSTGRESQL_CONNECTION"`
	FilePath              string `envconfig:"FILE_PATH"`
	WebServerAddr         string `envconfig:"WEB_SERVER_ADDR"`
	ExternalWebServerAddr string `envconfig:"EXTERNAL_WEB_SERVER_ADDR"`
	WebStaticDir          string `envconfig:"WEB_STATIC_DIR"`
	APIToken              string `envconfig:"API_TOKEN"`
	Debug                 bool   `envconfig:"DEBUG"`
}

func parseConfig() (Config, error) {
	c := Config{}

	err := envconfig.Process("APP", &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
