package main

import ()

type MessageRouter struct {
	Client *Client
}

func NewMessageRouter(client *Client) MessageRouter {
	router := MessageRouter{Client: client}
	return router
}

func (router *MessageRouter) Run() {
	context := NewContext(router.Client)
	for {
		select {
		case message := <-router.Client.ReaderChan:
			context.LatestMessage = message
			context = CmdUserJoin(context)
			context = CmdUserJoin(context)
			context = CmdPrivateMessasge(context)
			context = CmdMessageBroadcast(context)
			context = CmdJoinChannel(context)
			context = CmdLeaveChannel(context)

		case stop := <-router.Client.RoutingChan:
			if stop {
				return
			}
		}
	}
}
