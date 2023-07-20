package configuration

type BillingManagementConfig struct {
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Prefix string `yaml:"prefix"`
}

func GetBillingManagementConfig() *BillingManagementConfig {
	return &BillingManagementConfig{
		Host:   CS.GetString(string(BillingManagementConfigHost)),
		Port:   CS.GetString(string(BillingManagementConfigPort)),
		Prefix: CS.GetString(string(BillingManagementConfigPrefix)),
	}
}
