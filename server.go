package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	l       *net.Listener
	Host    string
	Port    int
	Clients []*Client
	Running bool
	Lock    sync.Mutex
}

func HandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func configTcp(l *net.Listener) {
}

func NewServer(host string, port int) *Server {
	server := &Server{Host: host, Port: port}
	server.Clients = make([]*Client, 100)
	return server
}

func (s *Server) Start() *Server {
	fmt.Println(fmt.Sprintf("%s:%d", s.Host, s.Port))
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port))
	HandleErr(err)
	for {
		conn, err := ln.Accept()
		HandleErr(err)
		s.Lock.Lock()
		client := NewClient(&conn)
		s.Clients = append(s.Clients, client)
		client.Loop()
		s.Lock.Unlock()
	}
	return s
}

func (s *Server) Stop() {
	for _, client := range s.Clients {
		client.Kick()
	}
	s.Running = false
}

func (s *Server) ClientLen() int {
	return len(s.Clients)
}
