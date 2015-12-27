package main

type Context struct {
	client        *Client
	RespMessages  []*Message
	LatestMessage *Message
}

func NewContext(client *Client) *Context {
	context := &Context{}
	return context
}

//need to flush all messages to client
func (context *Context) Flush() *Context {
	return context
}
