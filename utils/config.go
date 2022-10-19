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
	conf.Env = "prod"

	if os.Getenv("env") != "" {
		conf.Env = os.Getenv("env")
	}

	return conf
}
