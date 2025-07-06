package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

// env-default:"production"
type config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *config { // *config merupakan config struct
	var configPath string

	configPath = os.Getenv("CONFIG_PATH") //ambil config path

	if configPath == "" {
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse() // memberikan flags

		configPath = *flags // variabel ini isinya variabel flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file is not exist: %s", configPath)
	}

	var cfg config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("Can not read config file %s", err.Error())
	}

	return &cfg
}
