package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type Client struct {
	Conn       *net.Conn
	User       *User
	Reader     *bufio.Reader
	Writer     *bufio.Writer
	ReaderChan chan *Message
	WriterChan chan *Message
}

func NewClient(con *net.Conn) *Client {
	client := &Client{}
	client.Conn = con
	client.Reader = bufio.NewReader(*client.Conn)
	client.Writer = bufio.NewWriter(*client.Conn)
	return client
}

func (client *Client) SetUser(user *User) *Client {
	client.User = user
	return client
}

func (client *Client) Write() {
	for {
		select {
		case message := <-client.WriterChan:
			bytes, err := json.Marshal(message)
			HandleErr(err)
			client.Writer.Write(bytes)
		default:
		}
	}
}
func (client *Client) Read() {
	for {
		buf, err := client.Reader.ReadBytes('\n')
		HandleErr(err)
		message := Message{}
		json.Unmarshal(buf, &message)
		fmt.Println("GOT: ", message.MessageContent)
		//client.ReaderChan <- &message
	}
}

func (client *Client) Loop() {
	go client.Write()
	go client.Read()
}

func (client *Client) Kick() {
	fmt.Println("KicK")
}
