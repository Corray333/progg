package config

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

var Config config

type config struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

func GetAddress() string {
	return Config.Host + ":" + strconv.Itoa(Config.Port)
}

func LoadConfig() {
	cfg, err := os.ReadFile("../config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(cfg, &Config); err != nil {
		log.Fatal(err)
	}
	// TODO: setup logger
}
