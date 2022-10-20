package utils

import "os"

type Config struct {
	Env string
	Port int
	AppName string
}


func LoadConfig() Config {
	conf := Config{}

	conf.AppName = "TopMusicStreaming API"
	conf.Port = 9990
	conf.Env = PROD

	if os.Getenv("env") != EMPTY {
		conf.Env = os.Getenv("env")
	}

	return conf
}
