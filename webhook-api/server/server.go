package server

import (
	"database/sql"

	"github.com/aryanmaurya66/webhook/webhook-api/config"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
	db     *sql.DB
	cfg    *config.Config
}

func New(logger *logrus.Logger, db *sql.DB, cfg *config.Config) *Server {
	return &Server{logger: logger, db: db, cfg: cfg}
}

func (server *Server) RegisterAndStart() {

}
