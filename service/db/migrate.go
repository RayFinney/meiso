package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Migrate(conn *pgxpool.Pool) error {
	ctx := context.Background()
	if !hasMigrationTable(ctx, conn) {
		createMigrationTable(ctx, conn)
	}
	var migrated []string
	rows, err := conn.Query(ctx, "SELECT filename FROM migration")
	if err != nil {
		log.Error("unable to run migration: %v \n", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var migrateString string
		err := rows.Scan(&migrateString)
		if err != nil {
			log.Error("unable to run migration: %v \n", err)
			return err
		}
		migrated = append(migrated, migrateString)
	}
	var files []os.FileInfo
	var migratePath = "migrate/"
	err = filepath.Walk(migratePath, func(path string, info os.FileInfo, err error) error {
		files = append(files, info)
		return nil
	})
	if err != nil {
		return err
	}
	for _, file := range files {
		if file != nil && !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			if !isMigrated(migrated, file.Name()) {
				err = migrateFile(ctx, migratePath+file.Name(), file, conn)
				if err != nil {
					log.Error("unable to run migration (%s): %v \n", file.Name(), err)
					return err
				}
			}
		}
	}
	return nil
}

func hasMigrationTable(ctx context.Context, conn *pgxpool.Pool) bool {
	var dummy int64
	err := conn.QueryRow(ctx, "SELECT 1 FROM migration LIMIT 1").Scan(&dummy)
	if err != nil && err != pgx.ErrNoRows {
		return false
	}
	return true
}
func createMigrationTable(ctx context.Context, conn *pgxpool.Pool) {
	if err := os.Setenv("IMPORT", "TRUE"); err != nil {
		log.Infof("IMPORT ENV variable konnte nicht gesetzt werden, testdaten werden nicht importiert!")
	}

	_, err := conn.Exec(ctx, "CREATE TABLE migration ("+
		"id SERIAL PRIMARY KEY, "+
		"filename VARCHAR(255) NOT NULL"+
		");")
	if err != nil {
		panic(fmt.Sprintf("unable to create 'migration' table: %v", err))
	}
}
func isMigrated(migrated []string, filename string) bool {
	for _, migrate := range migrated {
		if migrate == filename {
			return true
		}
	}
	return false
}
func migrateFile(ctx context.Context, filepath string, info os.FileInfo, conn *pgxpool.Pool) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, string(data))
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, "INSERT INTO migration (filename) VALUES ($1);", info.Name())
	return err
}
