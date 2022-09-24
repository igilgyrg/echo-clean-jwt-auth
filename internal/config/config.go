package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

type Config struct {
	Port                  string `env:"PORT"`
	MongoPort             string `env:"MONGO_PORT"`
	MongoHost             string `env:"MONGO_HOST"`
	MongoDatabase         string `env:"MONGO_DATABASE"`
	MongoUsername         string `env:"MONGO_USERNAME"`
	MongoPassword         string `env:"MONGO_PASSWORD"`
	AccessTokenSignature  string `env:"ACCESS_TOKEN_SIGNATURE"`
	RefreshTokenSignature string `env:"REFRESH_TOKEN_SIGNATURE"`
}

var instance *Config
var once sync.Once

func NewConfig() *Config {
	once.Do(configEnv)
	return instance
}

func configEnv() {
	err := godotenv.Load()
	if err != nil {
		// TODO logger
		log.Fatal(err)
	}

	instance = &Config{}
	if err := env.Parse(instance); err != nil {
		// TODO logger
		log.Fatal(err)
	}
}
