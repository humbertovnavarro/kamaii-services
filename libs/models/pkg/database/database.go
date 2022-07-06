package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(createDSN()), &gorm.Config{})
}

func createDSN() string {
	sslmode := warnGetEnv("DB_SSL", "disabled")
	timezone := warnGetEnv("DB_TIMEZONE", "America/Los_Angeles")
	host := assertGetEnv("DB_HOST")
	user := assertGetEnv("DB_USER")
	password := assertGetEnv("DB_PASSWORD")
	dbname := assertGetEnv("DB_NAME")
	port := assertGetEnv("DB_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", host, user, password, dbname, port, sslmode, timezone)
}

func warnGetEnv(env string, def string) string {
	value := os.Getenv(env)
	if value == "" {
		fmt.Printf("missing environment variable %s, defaulting to %s\n", env, def)
		return def
	}
	return value
}

func assertGetEnv(env string) string {
	value := os.Getenv(env)
	if value == "" {
		panic(fmt.Sprintf("missing required environment variable: %s", env))
	}
	return value
}
