package tool

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type envExtraction struct {
	PORT         string
	APP_ENV      string
	DB_HOST      string
	DB_PORT      string
	DB_DATABASE  string
	DB_USERNAME  string
	DB_PASSWORD  string
	JWT_SECRET   string
	JWT_LIFETIME int
}

var EnvData *envExtraction

func NewEnv(path *string) *envExtraction {
	if EnvData != nil {
		return EnvData
	}
	if path != nil {
		err := godotenv.Load(*path)
		if err != nil {
			log.Fatal("tool_env_extraction.go|NewEnv with path|Error loading .env file")
		}
	}
	if path == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("tool_env_extraction.go|NewEnv without path|Error loading .env file")
		}
	}

	log.Default().Print("Init env object")

	jwt_lifetime := os.Getenv("JWT_LIFETIME")
	jwt_lifetime_int, _ := strconv.Atoi(jwt_lifetime)

	EnvData = &envExtraction{
		PORT:    os.Getenv("PORT"),
		APP_ENV: os.Getenv("APP_ENV"),

		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),

		JWT_SECRET:   os.Getenv("JWT_SECRET"),
		JWT_LIFETIME: jwt_lifetime_int,
	}
	log.Default().Print("Env data:", EnvData)
	return EnvData
}
