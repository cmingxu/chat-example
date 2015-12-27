package main

type MessasgeBroadcast struct {
	user *User
}

func (client *Client) CmdMessageBroadcast() *User {
	user := &User{}
	return user
}
