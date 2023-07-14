package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

// insert new user in users table in db
// returns error
func InsertUser(conn *pgx.Conn, email, nick, pass, pow string) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO users (username, password, pow, email)
		 VALUES ($1, $2, $3, $4)`,
		nick, pass, pow, email,
	)
	if err != nil {
		return err
	}
	return nil
}
