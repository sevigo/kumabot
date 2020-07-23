package ping

import (
	"strings"

	"github.com/sevigo/kumabot/pkg/core"
)

type pingBot struct {
	keyWord string
	name    string
}

func New() core.KumaBot {
	return &pingBot{
		keyWord: "ping",
		name:    "ping",
	}
}

func (p *pingBot) Match(in string) bool {
	in = strings.ToLower(in)
	in = strings.TrimSpace(in)
	return in == p.keyWord
}

func (p *pingBot) Run(_ string) string {
	return "pong"
}

func (p *pingBot) Name() string {
	return p.name
}
