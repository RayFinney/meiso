package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

func New() *pgxpool.Pool {
	var (
		user      = os.Getenv("SQL_USER")
		password  = os.Getenv("SQL_PASSWORD")
		address   = os.Getenv("SQL_ADDRESS")
		port      = os.Getenv("SQL_PORT")
		database  = os.Getenv("SQL_DATABASE")
		socketDir = os.Getenv("SQL_SOCKET_DIR")
	)
	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s%s port=%s", user, password, database, socketDir, address, port)
	conn, err := pgxpool.Connect(context.Background(), dbURI)
	if err != nil {
		panic(fmt.Sprintf("Verbindung mit Datenbank kann nicht hergestellt werden: %v", err))
	}
	return conn
}
