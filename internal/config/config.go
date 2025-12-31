package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTP_server struct {
	Address string
}

type Config struct {
	Env          string `yaml:"env" env-required:"true"` //env-default:"production"
	Storage_path string `yaml:"storage_path"`
	HTTP_server  `yaml:"http_server"`
}

func MustLoad() *Config{
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == ""{
		flags := flag.String("config","","path to configuration file")
		flag.Parse();

		configPath = *flags

		if configPath == ""{
			log.Fatal("config path not found")
		}

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err){
		log.Fatalf("Config path not found %s",configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath,&cfg)

	if err != nil{
		log.Fatalf("cannot read config file: %s",err.Error())
	}

	return &cfg
}