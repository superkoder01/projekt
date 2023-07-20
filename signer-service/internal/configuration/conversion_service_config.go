package configuration

import "time"

type ConversionServiceConfig struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Prefix  string        `yaml:"prefix"`
	Url     string        `yaml:"url"`
	Timeout time.Duration `yaml:"timeout"`
}

func GetConversionServiceConfigConfig() *ConversionServiceConfig {
	return &ConversionServiceConfig{
		Host:    CS.GetString(string(ConversionServiceConfigHost)),
		Port:    CS.GetString(string(ConversionServiceConfigPort)),
		Prefix:  CS.GetString(string(ConversionServiceConfigPrefix)),
		Url:     CS.GetString(string(ConversionServiceConfigUrl)),
		Timeout: CS.GetDuration(string(ConversionServiceConfigTimeout)),
	}
}
