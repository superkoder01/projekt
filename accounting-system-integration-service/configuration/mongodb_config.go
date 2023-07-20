package configuration

type MongoConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func GetMongoDatabaseConfig() *MongoConfig {
	cf := MongoConfig{
		User:     CS.GetString(string(MongoUser)),
		Password: CS.GetString(string(MongoPass)),
		Host:     CS.GetString(string(MongoHost)),
		Port:     CS.GetString(string(MongoPort)),
		Database: CS.GetString(string(MongoName)),
	}

	return &cf
}
