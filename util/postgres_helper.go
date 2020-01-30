package util

import (
	"database/sql"
	"fmt"

	"github.com/galihrivanto/dbutil"
	_ "github.com/lib/pq"
)

// PgDial initialize connection to postgres database
func PgDial(config dbutil.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=require",
		config.GetString("user", "postgres"),
		config.GetString("password", "postgres"),
		config.GetString("host", "localhost"),
		config.GetInt("port", 5432),
		config.GetString("dbname"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// add conn setting

	return db, nil
}
