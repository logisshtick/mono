package utils

import (
	"fmt"
	"os"
)

const (
	dbUrl = "postgres"
	noSsl = "sslmode=disable"
)

// get db url from env vars end constants
func GetDbUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:5432/postgres?%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		dbUrl,
		noSsl,
	)
}
