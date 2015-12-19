package main

import (
	"fmt"
)

func main() {
	fmt.Println("xx")
	server := NewServer("localhost", 10000)
	server.Start()

}
