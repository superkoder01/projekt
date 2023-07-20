package configuration

type authConfig struct {
	KeyPath string
}

func GetAuthConfig() *authConfig {
	cf := authConfig{
		KeyPath: CS.GetString(string(KeyPath)),
	}

	return &cf
}
