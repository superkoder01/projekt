package configuration

type skipJwtCheckConfig struct {
	Urls []string
}

func GetSkipJwtCheckConfig() *skipJwtCheckConfig {
	cf := skipJwtCheckConfig{
		Urls: CS.GetStringSlice(string(SkipJwtCheckUrls)),
	}
	return &cf
}
