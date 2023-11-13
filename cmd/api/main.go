package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

const version = "0.0.1"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	charmLogger := log.NewWithOptions(os.Stderr, log.Options{
		Prefix:       "http",
		ReportCaller: true,
	})

	app := &application{
		config: cfg,         // Configuration variables
		logger: charmLogger, // Colurful and human readable logger
	}

	// Wrap our app logger with standard logger for use in http.Server
	// Bease it only accepts standard logger interface
	stdLogger := app.logger.StandardLog(log.StandardLogOptions{
		ForceLevel: log.ErrorLevel,
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Minute,
		WriteTimeout: 30 * time.Minute,
		ErrorLog:     stdLogger,
	}

	app.logger.Info(fmt.Sprintf("Starting %s at %s", cfg.env, srv.Addr))

	err := srv.ListenAndServe()

	app.logger.Fatal(err)
}
