package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("doesn't exist: %s", configPath)
	}
	var cfg Config
	cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cn't read cfg file : %s", err.Error())
	}
	return &cfg
}
