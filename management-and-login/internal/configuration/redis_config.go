package configuration

type redisConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
}

func GetRedisConfig() *redisConfig {
	rc := redisConfig{
		Username: CS.GetString(string(RedisUser)),
		Password: CS.GetString(string(RedisPass)),
		Address:  CS.GetString(string(RedisAddress)),
	}

	return &rc
}
