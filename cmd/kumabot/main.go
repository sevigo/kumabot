package main

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/sevigo/kumabot/cmd/kumabot/config"
	"github.com/sevigo/kumabot/pkg/core"
	"github.com/sevigo/kumabot/pkg/server"
)

func main() {
	conf, err := config.Environ()
	if err != nil {
		log.Fatal("main: invalid configuration")
	}

	ctx := context.Background()

	app, err := InitializeApplication(ctx, conf)
	if err != nil {
		log.Fatal("main: cannot initialize server")
	}

	log.WithFields(log.Fields{
		"port": conf.Server.Port,
	}).Info("starting the http server")

	err = app.server.ListenAndServe(ctx)
	if err != nil {
		log.WithError(err).Error("Server was stopped with an error")
	} else {
		log.Print("Server was successfully stopped")
	}
}

// application is the main struct.
type application struct {
	server  *server.Server
	lineBot core.LineBot
}

func newApplication(srv *server.Server, lineBot core.LineBot) application {
	return application{
		server:  srv,
		lineBot: lineBot,
	}
}
