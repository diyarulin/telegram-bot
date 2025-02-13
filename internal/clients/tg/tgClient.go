package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"

	"github.com/diyarulin/telegram-bot/internal/model/messages"
	"github.com/pkg/errors"
)

type Client struct {
	Client *tgbotapi.BotAPI
}

type TokenGetter interface {
	Token() string
}

func New(tokenGetter TokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())
	if err != nil {
		return nil, errors.Wrap(err, "NewBotApi")
	}
	return &Client{
		Client: client,
	}, nil
}

func (c *Client) SendMessage(text string, userId int64) error {
	_, err := c.Client.Send(tgbotapi.NewMessage(userId, text))
	if err != nil {
		return errors.Wrap(err, "client.Send Error")
	}
	return nil
}

func (c *Client) ListenUpdates(msgModel *messages.Model) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.Client.GetUpdatesChan(u)

	log.Println("listening for messages")

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			err := msgModel.IncomingMessage(messages.Message{
				Text:   update.Message.Text,
				UserId: update.Message.From.ID,
			})
			if err != nil {
				log.Println("error processing message:", err)
			}
		}
	}
}
