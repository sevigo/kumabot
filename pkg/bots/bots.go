package bots

import (
	"context"

	"github.com/sevigo/kumabot/pkg/bots/ping"
	"github.com/sevigo/kumabot/pkg/bots/weather"
	"github.com/sevigo/kumabot/pkg/core"
)

type Bots struct {
	ctx context.Context
}

var botsCatalog = []core.KumaBot{
	ping.New(),
	weather.New(),
}

func New(ctx context.Context) core.KumaBots {
	return &Bots{
		ctx: ctx,
	}
}

func (b *Bots) All() []core.KumaBot {
	return botsCatalog
}
