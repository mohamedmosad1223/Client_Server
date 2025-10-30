package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Sender  string
	Content string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Could not connect to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	// اسم المستخدم
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Type messages to send. Type 'exit' to quit.")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting...")
			break
		}

		msg := Message{Sender: name, Content: text}
		var reply []string
		err = client.Call("ChatServer.SendMessage", msg, &reply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		// عرض التاريخ كله
		fmt.Println("\n--- Chat History ---")
		for _, m := range reply {
			fmt.Println(m)
		}
		fmt.Println("-------------------")
	}
}
