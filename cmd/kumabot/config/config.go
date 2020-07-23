package config

import (
	"net/http"
	"os"
)

type Config struct {
	LineBot LineBot
	Logging Logging
	Server  Server
}

type Logging struct {
	Debug  bool
	Trace  bool
	Color  bool
	Pretty bool
	Text   bool
}

type Server struct {
	Addr    string
	Host    string
	Port    string
	Proto   string
	Handler http.Handler
}

type LineBot struct {
	ChannelSecret string
	ChannelToken  string
}

func Environ() (Config, error) {
	cfg := Config{}
	defaultAddress(&cfg)
	defaultLogger(&cfg)
	lineConfigFromEnv(&cfg)
	return cfg, nil
}

func defaultLogger(c *Config) {
	c.Logging = Logging{
		Debug: true,
		Color: true,
		Text:  true,
	}
}

func defaultAddress(c *Config) {
	if c.Server.Host != "" && c.Server.Proto != "" && c.Server.Port != "" {
		c.Server.Addr = c.Server.Proto + "://" + c.Server.Host + ":" + c.Server.Port
	} else {
		c.Server.Port = "8081"
		c.Server.Addr = ":" + c.Server.Port
	}
}

func lineConfigFromEnv(c *Config) {
	c.LineBot.ChannelSecret = os.Getenv("LINE_CHANNEL_SECRET")
	c.LineBot.ChannelToken = os.Getenv("LINE_CHANNEL_TOKEN")
}
