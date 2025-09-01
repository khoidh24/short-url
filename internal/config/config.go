package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBUrl      string
	Port       string
	SecretKey  string
	ApiKey     string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("PGUSER")
	password := url.QueryEscape(os.Getenv("PGPASSWORD"))
	host := os.Getenv("PGHOST")
	dbname := os.Getenv("PGDATABASE")

	dbUrl := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=require&search_path=public",
		user, password, host, dbname,
	)

	return Config{
		DBUser:     user,
		DBPassword: password,
		DBHost:     host,
		DBUrl:      dbUrl,
		Port:       os.Getenv("PORT"),
		SecretKey:  os.Getenv("SECRET_KEY"),
		ApiKey:     os.Getenv("API_KEY"),
	}
}
