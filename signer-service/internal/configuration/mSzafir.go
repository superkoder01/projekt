package configuration

type MSzafirConfig struct {
	Host                            string `yaml:"host"`
	Port                            string `yaml:"port"`
	Prefix                          string `yaml:"prefix"`
	Timestamp                       int64  `yaml:"timestamp"`
	Mode                            string `yaml:"mode"`
	UrlSigningCompleted             string `yaml:"urlSigningCompleted"`
	UrlSigningCompletedNotification string `yaml:"urlSigningCompletedNotification"`
	Password                        string `yaml:"password"`
	SignedContractsPath             string `yaml:"signedContractsPath"`
}

func GetMSzafirConfig() *MSzafirConfig {
	return &MSzafirConfig{
		Host:                            CS.GetString(string(MSzafirHost)),
		Port:                            CS.GetString(string(MSzafirPort)),
		Prefix:                          CS.GetString(string(MSzafirPrefix)),
		Timestamp:                       int64(CS.GetInt(string(MSzafirTimestamp))),
		Mode:                            CS.GetString(string(MSzafirMode)),
		UrlSigningCompleted:             CS.GetString(string(MSzafirUrlSigningCompleted)),
		UrlSigningCompletedNotification: CS.GetString(string(MSzafirUrlSigningCompletedNotification)),
		Password:                        CS.GetString(string(MSzafirPassword)),
		SignedContractsPath:             CS.GetString(string(MSzafirSignedContractsPath)),
	}
}
