package configuration

type Logger struct {
	Development       bool   `yaml:"Development"`
	DisableStacktrace bool   `yaml:"DisableStacktrace"`
	Encoding          string `yaml:"Encoding"`
	Level             string `yaml:"Level"`
}

func GetLoggerConfig() *Logger {
	return &Logger{
		Development:       CS.GetBool(string(LoggerDevelopment)),
		DisableStacktrace: CS.GetBool(string(LoggerDisableStacktrace)),
		Encoding:          CS.GetString(string(LoggerEncoding)),
		Level:             CS.GetString(string(LoggerLevel)),
	}
}
