package webhook

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"

	"github.com/sevigo/kumabot/pkg/core"
)

type BotServer struct {
	Logger   *logrus.Logger
	LineBot  core.LineBot
	KumaBots core.KumaBots
}

func New(lineBot core.LineBot, kumaBots core.KumaBots, logger *logrus.Logger) *BotServer {
	return &BotServer{
		LineBot:  lineBot,
		KumaBots: kumaBots,
		Logger:   logger,
	}
}

func (s *BotServer) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Post("/", HandleEcho(s.LineBot, s.KumaBots, s.Logger))

	return r
}

func HandleEcho(bot core.LineBot, kumaBots core.KumaBots, log *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// <debug http>
		// requestDump, err := httputil.DumpRequest(req, true)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(requestDump))
		// </debug http>

		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			log.WithError(err).Error("linebot.ParseRequest() error")
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				for _, kumaBot := range kumaBots.All() {
					switch message := event.Message.(type) {
					case *linebot.TextMessage:
						if kumaBot.Match(message.Text) {
							text := kumaBot.Run(message.Text)
							log.WithField("bot", kumaBot.Name()).
								Infof("response: %q", text)

							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(text)); err != nil {
								log.WithError(err).
									WithField("bot", kumaBot.Name).
									Error("ReplyMessage() error")
							}
						}
					}
				}
			}
		}
	}
}
