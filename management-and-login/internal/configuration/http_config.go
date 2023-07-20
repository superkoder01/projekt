package configuration

type httpConfig struct {
	Port      string `yaml:"port"`
	ApiPrefix string `yaml:"apiPrefix"`
}

func GetHttpConfig() *httpConfig {
	hf := httpConfig{
		Port:      CS.GetString(string(ListenPort)),
		ApiPrefix: CS.GetString(string(ApiPrefix)),
	}

	return &hf
}
