package main

import (
	"fmt"
)

type MessageRouter struct {
	Client *Client
}

func NewMessageRouter(client *Client) MessageRouter {
	router := MessageRouter{Client: client}
	return router
}

func (router *MessageRouter) Run() {
	for {
		select {
		case message := <-router.Client.ReaderChan:
			fmt.Println(message.MessageContent)
			echoMessage := Message{MessageContent: fmt.Sprintf("echo: %s", message.MessageContent)}
			router.Client.WriterChan <- &echoMessage
		case stop := <-router.Client.RoutingChan:
			if stop {
				return
			}
		}
	}
}