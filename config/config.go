package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
	SeverPort  string
	GeoKey     string
}

func InitConfig() *Config {
	var result = new(Config)
	result = loadConfig()

	if result == nil {
		log.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return result
}
func loadConfig() *Config {
	var result = new(Config)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Config: Cannot load config file,", err.Error())
		return nil
	}

	if value, found := os.LookupEnv("DB_USER"); found {
		result.DBUsername = value
	}
	if value, found := os.LookupEnv("DB_PASS"); found {
		result.DBPassword = value
	}
	if value, found := os.LookupEnv("DB_PORT"); found {
		result.DBPort = value
	}
	if value, found := os.LookupEnv("DB_HOST"); found {
		result.DBHost = value
	}
	if value, found := os.LookupEnv("DB_NAME"); found {
		result.DBName = value
	}
	if value, found := os.LookupEnv("SERVER_PORT"); found {
		result.SeverPort = value
	}
	if value, found := os.LookupEnv("GEOAPIFY_KEY"); found {
		result.GeoKey = value
	}

	return result
}
