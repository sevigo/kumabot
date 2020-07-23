package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"

	"github.com/sevigo/kumabot/cmd/kumabot/config"
	"github.com/sevigo/kumabot/pkg/bots"
	"github.com/sevigo/kumabot/pkg/core"
	"github.com/sevigo/kumabot/pkg/server"

	"github.com/sevigo/kumabot/pkg/handler/health"
	"github.com/sevigo/kumabot/pkg/handler/webhook"
)

type healthzHandler http.Handler

var serverSet = wire.NewSet(
	provideServer,
	provideHealthz,
	provideLogger,
	provideRouter,
	provideKumaBots,
	webhookProvider,
)

func provideServer(handler http.Handler, config config.Config) *server.Server {
	return &server.Server{
		Addr:    config.Server.Addr,
		Host:    config.Server.Host,
		Handler: handler,
	}
}

func provideLogger(config config.Config) *logrus.Logger {
	logger := logrus.StandardLogger()
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: config.Logging.Color,
	})
	logger.SetReportCaller(false)
	if config.Logging.Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
	logger.SetOutput(os.Stdout)
	return logger
}

func provideKumaBots(ctx context.Context) core.KumaBots {
	return bots.New(ctx)
}

func webhookProvider(lineBot core.LineBot, kumaBots core.KumaBots, logger *logrus.Logger) *webhook.BotServer {
	return webhook.New(lineBot, kumaBots, logger)
}

func provideRouter(healthz healthzHandler, webhook *webhook.BotServer) http.Handler {
	r := chi.NewRouter()
	r.Mount("/healthz", healthz)
	r.Mount("/", webhook.Handler())
	return r
}

func provideHealthz() healthzHandler {
	v := health.New()
	return healthzHandler(v)
}
