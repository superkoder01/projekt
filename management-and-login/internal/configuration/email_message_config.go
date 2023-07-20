package configuration

type emailMessage struct {
	Title string `yaml:"title"`
	Body  string `yaml:"body"`
}

func GetActivationMessageConfig() *emailMessage {
	cf := emailMessage{
		Title: CS.GetString(string(EmailActivationTitle)),
		Body:  CS.GetString(string(EmailActivationBody)),
	}

	return &cf
}

func GetPasswordResetMessageConfig() *emailMessage {
	cf := emailMessage{
		Title: CS.GetString(string(EmailResetPasswordTitle)),
		Body:  CS.GetString(string(EmailResetPasswordBody)),
	}

	return &cf
}
