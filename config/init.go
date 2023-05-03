package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	DatabaseURL string `env:"DATABASE_URL"`
}

var Cfg Config

func Init() {
	err := cleanenv.ReadConfig(".env", &Cfg)
	if err != nil {
		panic(err)
	}
}
