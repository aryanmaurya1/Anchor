package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func SetupDatabase(logger *logrus.Logger, filename string) (*sql.DB, error) {
	db, err := getDatabaseConnection(logger, filename)
	if err != nil {
		return nil, err
	}

	err = migrate(logger, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDatabaseConnection(logger *logrus.Logger, filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		logger.Errorf("[ERROR GETTING DB CONNECTION] %+v \n", err.Error())
		return nil, err
	}
	return db, nil
}

func migrate(logger *logrus.Logger, db *sql.DB) error {
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS "buckets" (
			"id"	TEXT NOT NULL UNIQUE,
			"bucket_name"	TEXT NOT NULL UNIQUE,
			"created_at"	INTEGER NOT NULL,
			PRIMARY KEY("id")
		);
		CREATE TABLE IF NOT EXISTS "records" (
			"id"	TEXT NOT NULL UNIQUE,
			"bucket_id"	TEXT NOT NULL,
			"headers"	TEXT,
			"body"	TEXT,
			"created_at"	INTEGER NOT NULL,
			FOREIGN KEY("bucket_id") REFERENCES "buckets"("id"),
			PRIMARY KEY("id")
		);
	;`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		logger.Errorf("[ERROR MIGRATING DATABASE] %+v \n", err.Error())
		return err
	}
	return nil
}
