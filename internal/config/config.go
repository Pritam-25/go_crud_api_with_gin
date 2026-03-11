package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI    string
	MongoDBName string
	Port        string
}

func LoadConfig() (*Config, error) {
	// load .env file (ignore error in production)
	_ = godotenv.Load()

	mongoURI, err := getEnv("MONGO_URI")
	if err != nil {
		return nil, err
	}

	mongoDBName, err := getEnv("MONGO_DB_NAME")
	if err != nil {
		return nil, err
	}

	port, err := getEnv("PORT")
	if err != nil {
		return nil, err
	}

	return &Config{
		MongoURI:    mongoURI,
		MongoDBName: mongoDBName,
		Port:        port,
	}, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)

	if value == "" {
		return "", fmt.Errorf("%s environment variable not set", key)
	}

	return value, nil
}