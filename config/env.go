package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	PostgresDBHost     string
	PostgresDBPort     string
	PostgresDBUser     string
	PostgresDBPassword string
	PostgresDBName     string
	PostgresDBSSLMode  string
	GormMode           string
	GinMode            string
	ServerPort         string
	log                ILogger
}

var (
	envInstance *Env
	envonce     sync.Once
)

func NewGetEnv() *Env {
	envonce.Do(func() {
		log := NewLogger()
		err := godotenv.Load("./.env") // Ajusta la ruta relativa a tu archivo .env
		if err != nil {
			err = godotenv.Load("/app/.env")
		}

		if err != nil {
			log.Fatal("Error al cargar el archivo .env: %v", err)
		}

		envInstance = &Env{
			PostgresDBHost:     os.Getenv("POSTGRES_DB_HOST"),
			PostgresDBPort:     os.Getenv("POSTGRES_DB_PORT"),
			PostgresDBUser:     os.Getenv("POSTGRES_DB_USER"),
			PostgresDBPassword: os.Getenv("POSTGRES_DB_PASSWORD"),
			PostgresDBName:     os.Getenv("POSTGRES_DB_NAME"),
			PostgresDBSSLMode:  os.Getenv("POSTGRES_DB_SSLMODE"),
			GormMode:           os.Getenv("GORM_MODE"),
			GinMode:            os.Getenv("GIN_MODE"),
			ServerPort:         os.Getenv("SERVER_PORT"),
			log:                NewLogger(),
		}
	})

	return envInstance
}

func validateEnvVariables(env *Env) {
	if env.PostgresDBHost == "" {
		log.Fatal("POSTGRES_DB_HOST is required but not set")
	}
	if env.PostgresDBPort == "" {
		log.Fatal("POSTGRES_DB_PORT is required but not set")
	}
	if env.PostgresDBUser == "" {
		log.Fatal("POSTGRES_DB_USER is required but not set")
	}
	if env.PostgresDBPassword == "" {
		log.Fatal("POSTGRES_DB_PASSWORD is required but not set")
	}
	if env.PostgresDBName == "" {
		log.Fatal("POSTGRES_DB_NAME is required but not set")
	}
	if env.ServerPort == "" {
		log.Fatal("SERVER_PORT is required but not set")
	}
	if env.GormMode == "" {
		log.Fatal("GORM_MODE is required but not set")
	}
	if env.GinMode == "" {
		log.Fatal("GIN_MODE is required but not set")
	}

}
