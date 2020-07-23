package webhook

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sevigo/kumabot/mocks"
	"github.com/sevigo/kumabot/pkg/core"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandleHealthz(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	botMessageText := "test"
	responseText := "test-test"
	replyToken := "11223344"

	lineBot := mocks.NewMockLineBot(controller)
	lineBot.EXPECT().ParseRequest(gomock.Any()).DoAndReturn(func(req *http.Request) ([]*linebot.Event, error) {
		return []*linebot.Event{
			{
				Type:       linebot.EventTypeMessage,
				Message:    linebot.NewTextMessage(botMessageText),
				ReplyToken: replyToken,
			},
		}, nil
	})
	lineBot.EXPECT().ReplyMessage(replyToken, linebot.NewTextMessage(responseText)).Return(&linebot.BasicResponse{
		RequestID: "123",
	}, nil)

	kumaBot := mocks.NewMockKumaBot(controller)
	kumaBot.EXPECT().Match(botMessageText).Return(true)
	kumaBot.EXPECT().Run(botMessageText).Return(responseText)
	kumaBot.EXPECT().Name().Return("test")

	kumaBots := mocks.NewMockKumaBots(controller)
	kumaBots.EXPECT().All().Return([]core.KumaBot{kumaBot})

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBufferString(""))

	logger := logrus.StandardLogger()
	New(lineBot, kumaBots, logger).Handler().ServeHTTP(w, r)
	assert.Equal(t, 200, w.Code)
}
