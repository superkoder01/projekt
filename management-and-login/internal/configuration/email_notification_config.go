package configuration

type emailConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Queue    string `yaml:"queue"`
}

func GetEmailNotificationConfig() *emailConfig {
	cf := emailConfig{
		User:     CS.GetString(string(EmailUser)),
		Password: CS.GetString(string(EmailPassword)),
		Host:     CS.GetString(string(EmailHost)),
		Port:     CS.GetString(string(EmailPort)),
		Queue:    CS.GetString(string(EmailQueue)),
	}

	return &cf
}
