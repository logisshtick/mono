package utils

import (
	"fmt"
	"os"
)

const (
	dbUrl = "postgres"
)

// get db url from env vars end constants
func GetDbUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:5432/postgres?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		dbUrl,
	)
}
