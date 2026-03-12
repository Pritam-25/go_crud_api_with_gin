package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	MongoURI    string `env:"MONGO_URI" env-required:"true"`
	MongoDBName string `env:"MONGO_DB_NAME" env-required:"true"`
	Port        string `env:"PORT" env-required:"true"`
	JWTSecret   string `env:"JWT_SECRET" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
