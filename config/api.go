package config

type API struct {
	Addr         string `yaml:"addr" envconfig:"ADDR"`
	ExternalAddr string `yaml:"external_addr" envconfig:"EXTERNAL_ADDR"`
	StaticDir    string `yaml:"static_dir" envconfig:"STATIC_DIR"`
	Token        string `yaml:"token" envconfig:"TOKEN"`
}

func (a API) GetAddr() string {
	return a.Addr
}

func (a API) GetExternalAddr() string {
	return a.ExternalAddr
}

func (a API) GetStaticDir() string {
	return a.StaticDir
}

func (a API) GetToken() string {
	return a.Token
}

func APIDefault() API {
	return API{
		Addr:         ":8080",
		ExternalAddr: "http://localhost:8080",
		StaticDir:    "",
		Token:        "",
	}
}
