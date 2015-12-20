package main

const (
	MessagePlain = 1
)

type Message struct {
	FromId         int64 `json:"FromId"`
	ToId           int64 `json:"ToId"`
	Channel        *Channel
	ChannelId      int64  `json:"ChannelId"`
	MessageType    int    `json:"MessageType"`
	MessageContent string `json:"MessageContent"`
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
