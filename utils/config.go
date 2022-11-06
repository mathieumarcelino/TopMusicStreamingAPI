package utils

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Env     string
	Port    int
	AppName string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		Logger.Fatal("could not load .env file")
	}

	conf := Config{}

	conf.AppName = "TopMusicStreaming API"
	conf.Port = 9990
	conf.Env = PROD

	if os.Getenv("ENV") != EMPTY {
		conf.Env = os.Getenv("ENV")
	}

	return conf
}
