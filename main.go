package main

import (
	"fmt"
	"path"

	"github.com/aryanmaurya66/webhook/webhook-api/config"
	"github.com/aryanmaurya66/webhook/webhook-api/pkg/sqlite"
	"github.com/aryanmaurya66/webhook/webhook-api/server"
	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	cfg := config.New(logger)
	config, err := cfg.ReadConfig(path.Join("webhook-api", "config", "config.json"))
	if err != nil {
		panic(err)
	}

	loglevel := logrus.DebugLevel
	loglevel, err = logrus.ParseLevel(config.Loglevel)
	if err != nil {
		fmt.Printf("[ERROR PARSING LOG LEVEL] %+v, [DEBUG]", err.Error())
	}
	logger.SetLevel(loglevel)

	db, err := sqlite.SetupDatabase(logger, config.Database.Filename)
	if err != nil {
		panic(err)
	}

	server := server.New(logger, db, config)
	server.RegisterAndStart()
}
