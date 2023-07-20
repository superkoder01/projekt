package configuration

type authConfig struct {
	KeyPath               string `yaml:"keyPath"`
	AccessExpirationTime  int    `yaml:"accessExpirationTime"`
	RefreshExpirationTime int    `yaml:"refreshExpirationTime"`
}

func GetAuthConfig() *authConfig {
	cf := authConfig{
		KeyPath:               CS.GetString(string(KeyPath)),
		AccessExpirationTime:  CS.GetInt(string(AccessExpirationTime)),
		RefreshExpirationTime: CS.GetInt(string(RefreshExpirationTime)),
	}

	return &cf
}
