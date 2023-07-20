package configuration

type blockchainConfig struct {
	AdapterHost string `yaml:"adapterHost"`
	AdapterPort string `yaml:"adapterPort"`
	Endpoint    string `yaml:"endpoint"`
}

func GetBlockchainConfig() *blockchainConfig {
	bf := blockchainConfig{
		AdapterHost: CS.GetString(string(AdapterHost)),
		AdapterPort: CS.GetString(string(AdapterPort)),
		Endpoint:    CS.GetString(string(AdapterEndpoint)),
	}

	return &bf
}
