package types

type MessageStore interface {
	GetMessage(id string) (Message, error)
	GetMessages() []Message
	CreateMessage(message Message) error
	UpdateMessage(id string, message *Message) error
	DeleteMessage(id string) error
}
