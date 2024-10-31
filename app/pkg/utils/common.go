package utils

import (
	"gopkg.in/yaml.v2"
	"log"
	"main/app/internal/config"
	"os"
)

func LoadConfig(path string) *config.Config {
	confStream, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("Error to open read config file:", err)
	}

	conf := config.New()
	err = yaml.Unmarshal(confStream, conf)
	if err != nil {
		log.Fatalln("Error to unmarshal data from config file:", err)
	}
	return conf
}
