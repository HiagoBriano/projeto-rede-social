package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringDatabaseConnection string
	Port                     int
)

// Load vai inicializar as variáveis de ambiente
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 3001
	}

	StringDatabaseConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
	)
}
