package messages

type MessageSender interface {
	SendMessage(text string, userId int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text   string
	UserId int64
}

func (s *Model) IncomingMessage(msg Message) error {
	if msg.Text == "/start" {

		s.tgClient.SendMessage("hello", msg.UserId)
		return nil
	}
	s.tgClient.SendMessage("I don't know this command", msg.UserId)
	return nil
}
