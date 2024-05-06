package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Site is the data structure that represents an Site.
type Environment struct {
	Host   string
	Port   int
	DbHost string
	DbUser string
	DbPass string
	DbPort int
	DbName string
}

// LoadEnvironment creates a new Environment.
func LoadEnvironment() Environment {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	godotenv.Load(".env")

	// getting env variables SITE_TITLE and DB_HOST
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	CheckConversionError("PORT", err)

	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbPort, err := strconv.Atoi(os.Getenv("DBPORT"))
	CheckConversionError("DBPORT", err)
	dbName := os.Getenv("DBNAME")

	return Environment{
		Host:   host,
		Port:   port,
		DbHost: dbHost,
		DbUser: dbUser,
		DbPass: dbPass,
		DbPort: dbPort,
		DbName: dbName,
	}
}

func CheckConversionError(varName string, err error) {
	if err != nil {
		log.Fatalf("%s environment variable not found", varName)
	}
}
