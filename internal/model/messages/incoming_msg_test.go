package messages

import (
	"testing"

	mocks "github.com/diyarulin/telegram-bot/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)
	sender.EXPECT().SendMessage("hello", int64(123))
	err := model.IncomingMessage(Message{
		Text:   "/start",
		UserId: 123,
	})

	assert.NoError(t, err)

}
func Test_OnUnknownCommand_ShouldAnswerWithHelpMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)
	sender.EXPECT().SendMessage("I don't know this command", int64(123))
	err := model.IncomingMessage(Message{
		Text:   "some text",
		UserId: 123,
	})

	assert.NoError(t, err)

}
