//+build wireinject

package main

import (
	"context"

	"github.com/google/wire"

	"github.com/sevigo/kumabot/cmd/kumabot/config"
)

func InitializeApplication(ctx context.Context, config config.Config) (application, error) {
	wire.Build(
		serverSet,
		lineBotSet,
		newApplication,
	)

	return application{}, nil
}
