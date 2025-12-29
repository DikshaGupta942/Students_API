package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Httpserver struct {
	Address string `yaml:"address" env_default:":8082"`
}

type Config struct {
	Env         string     `yaml:"env"`          //env:"ENV" env_required:"true" env_default:"production"`
	Storagepath string     `yaml:"storage_path"` // env_required:"true"`
	Httpserver  Httpserver `yaml:"http_server"`
}

func MustLoad() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	configPath := "config/" + env + ".yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %s", err.Error())
	}

	return &cfg
}

// func MustLoad() *Config {
// 	var configPath string
// 	configPath = os.Getenv("CONFIG_PATH")
// 	if configPath == "" {
// 		configPathFlag := flag.String("config", "", "Path to config file")

// 		flag.Parse()

// 		configPath = *configPathFlag

// 		if configPath == "" {
// 			log.Fatal("config path is required")
// 			//configPath = "config/config.yaml"
// 		}
// 	}

// 	if _, err := os.Stat(configPath); os.IsNotExist(err) {
// 		log.Fatalf("config file does not exist: %s", configPath)
// 	}

// 	var cfg Config

// 	err := cleanenv.ReadConfig(configPath, &cfg)
// 	if err != nil {
// 		log.Fatalf("failed to read config: %s", err.Error())
// 	}
// 	return &cfg

// }
