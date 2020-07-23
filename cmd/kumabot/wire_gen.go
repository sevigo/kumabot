// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"context"
	"github.com/sevigo/kumabot/cmd/kumabot/config"
)

// Injectors from wire.go:

func InitializeApplication(ctx context.Context, config2 config.Config) (application, error) {
	mainHealthzHandler := provideHealthz()
	lineBot := provideLineBot(config2)
	kumaBots := provideKumaBots(ctx)
	logger := provideLogger(config2)
	botServer := webhookProvider(lineBot, kumaBots, logger)
	handler := provideRouter(mainHealthzHandler, botServer)
	server := provideServer(handler, config2)
	mainApplication := newApplication(server, lineBot)
	return mainApplication, nil
}
