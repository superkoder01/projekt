package configuration

type dbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func GetDatabaseConfig() *dbConfig {
	cf := dbConfig{
		User:     CS.GetString(string(DbUser)),
		Password: CS.GetString(string(DbPass)),
		Host:     CS.GetString(string(DbHost)),
		Port:     CS.GetString(string(DbPort)),
		Database: CS.GetString(string(DbName)),
	}

	return &cf
}
