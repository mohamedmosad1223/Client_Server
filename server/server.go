package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

type Message struct {
	Sender  string
	Content string
}

func (c *ChatServer) SendMessage(msg Message, reply *[]string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.messages = append(c.messages, fmt.Sprintf("%s: %s", msg.Sender, msg.Content))

	*reply = c.messages
	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Chat server is running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
