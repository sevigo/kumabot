package core

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

type LineBot interface {
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
	ReplyMessage(replyToken string, messages ...linebot.SendingMessage) (*linebot.BasicResponse, error)
}

func NewLineBot(channelSecret, channelToken string) (LineBot, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}
	return &lineBotWrapper{
		Client: bot,
	}, nil
}

type lineBotWrapper struct {
	*linebot.Client
}

func (w *lineBotWrapper) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return w.Client.ParseRequest(r)
}

func (w *lineBotWrapper) ReplyMessage(replyToken string, messages ...linebot.SendingMessage) (*linebot.BasicResponse, error) {
	// TODO: add WithContext and maybe some retry logic, if this is not implemented in the client
	// return w.Client.ReplyMessage(replyToken, messages...).WithContext(context.TODO()).Do()
	return w.Client.ReplyMessage(replyToken, messages...).Do()
}
