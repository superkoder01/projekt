package configuration

type listenPortConfig struct {
	Port string
}

func GetListenPortConfig() *listenPortConfig {
	cf := listenPortConfig{
		Port: CS.GetString(string(ListenPort)),
	}

	return &cf
}
