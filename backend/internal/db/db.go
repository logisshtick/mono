package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

// insert new user in users table in db
// returns error
func InsertUser(conn *pgx.Conn, email, nick, pass string) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO users (username, password, email)
		 VALUES ($1, $2, $3)`,
		nick, pass, email,
	)
	if err != nil {
		return err
	}
	return nil
}
