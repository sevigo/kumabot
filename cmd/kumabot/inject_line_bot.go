package main

import (
	"github.com/google/wire"

	"github.com/sevigo/kumabot/cmd/kumabot/config"
	"github.com/sevigo/kumabot/pkg/core"
)

var lineBotSet = wire.NewSet(
	provideLineBot,
)

func provideLineBot(config config.Config) core.LineBot {
	bot, err := core.NewLineBot(config.LineBot.ChannelSecret, config.LineBot.ChannelToken)
	if err != nil {
		panic(err)
	}

	return bot
}
