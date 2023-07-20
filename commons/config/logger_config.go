package config

//LoggerConfig default Time encoder format(2006-01-02 15:04:05)
type LoggerConfig struct {
	Development       bool
	Encoding          string
	Level             string
	TimeEncoderFormat string
}
