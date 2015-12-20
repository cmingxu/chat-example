package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Client struct {
	Conn          *net.Conn
	User          *User
	Reader        *bufio.Reader
	Writer        *bufio.Writer
	ReaderChan    chan *Message
	WriterChan    chan *Message
	MessageRouter MessageRouter

	RoutingChan chan bool
	Writing     chan bool
	Reading     chan bool
}

func NewClient(con *net.Conn) *Client {
	client := &Client{}
	client.Conn = con
	client.Reader = bufio.NewReader(*client.Conn)
	client.Writer = bufio.NewWriter(*client.Conn)
	client.ReaderChan = make(chan *Message)
	client.WriterChan = make(chan *Message)
	client.RoutingChan = make(chan bool)
	client.Writing = make(chan bool)
	client.Reading = make(chan bool)
	client.MessageRouter = NewMessageRouter(client)
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
			fmt.Println("Writing: ", string(bytes))
			client.Writer.Write(bytes)
			client.Writer.Write([]byte("\n"))
			client.Writer.Flush()
		case stop := <-client.Writing:
			if stop {
				return
			}
		}
	}
}
func (client *Client) Read() {
	for {
		buf, err := client.Reader.ReadBytes('\n')
		if err != nil && err == io.EOF {
			client.GracefullyShutDown()
		}
		message := Message{}
		json.Unmarshal(buf, &message)
		client.ReaderChan <- &message
		fmt.Println("GOT: ", message.MessageContent)
	}
}

func (client *Client) Routing() {
	client.MessageRouter.Run()
}

func (client *Client) Loop() {
	go client.Routing()
	go client.Write()
	go client.Read()
}

func (client *Client) Kick() {
	fmt.Println("KicK")
}

func (client *Client) GracefullyShutDown() {
	fmt.Println("GracefullyShutDown")
	client.RoutingChan <- false
	client.Writing <- false
	client.Reading <- false
}
