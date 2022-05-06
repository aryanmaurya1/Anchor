package database

import (
	"database/sql"

	"github.com/sirupsen/logrus"
)

type DatabaseClient struct {
	logger *logrus.Logger
	db     *sql.DB
}
