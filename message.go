package main

const (
	MessagePlain = 1
)

type Message struct {
	FromId         int64
	ToId           int64
	Channel        *Channel
	ChannelId      int64
	MessageType    int
	MessageContent string
}

func NewMessage() *Message {
	return &Message{}
}

func SetFromId(m *Message, from_id int64) *Message {
	m.FromId = from_id
	return m
}

func SetToId(m *Message, to_id int64) *Message {
	m.ToId = to_id
	return m
}
