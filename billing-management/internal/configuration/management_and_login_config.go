package configuration

type ManagementAndLoginConfig struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Prefix string `yaml:"prefix"`
}

func GetManagementAndLoginConfig() *ManagementAndLoginConfig {
	return &ManagementAndLoginConfig{
		Host:   CS.GetString(string(ManagementAndLoginHost)),
		Port:   CS.GetString(string(ManagementAndLoginPort)),
		Prefix: CS.GetString(string(ManagementAndLoginPrefix)),
	}
}
