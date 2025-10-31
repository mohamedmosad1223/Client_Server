Go RPC Chat Project
A simple client-server chat application built in Go using RPC (Remote Procedure Call). This project demonstrates a basic chat system where multiple clients can connect to a server, send messages, and view the chat history.

🚀 Features
Connect multiple clients to a single server.

Send and receive messages in real-time.

View the full chat history after each message is sent.

Thread-safe message handling on the server using sync.Mutex.

📁 Project Structure
client-server/
├── client/
│   └── client.go
└── server/
    └── server.go
client/: Contains the client-side Go program that connects to the server.

server/: Contains the server-side Go program that manages the chat.

🏁 Getting Started
Prerequisites
Go (version 1.18+ recommended)

A basic understanding of Go programming.

🛠️ Installation & Running
Clone the repository.

Run the Server Open a terminal and navigate to the server directory:

Bash

cd server
go run server.go
You will see the following message, indicating the server is ready:

Chat server is running on port 1234...
Run the Client Open a new terminal window and navigate to the client directory:

Bash

cd client
go run client.go
Start Chatting!

You will be prompted to enter your name.

Type your messages and press Enter to send.

Type exit to quit.

Example Usage:

Enter your name: Alice
Type messages to send. Type 'exit' to quit.
> Hello, everyone!
--- Chat History ---
Alice: Hello, everyone!
-------------------
> exit
Exiting...
🔧 How It Works
Server
Listens on TCP port 1234.

Maintains a slice of all messages (messages []string).

Uses sync.Mutex to ensure thread-safety when multiple clients send messages concurrently.

Exposes a SendMessage method via RPC. This method accepts a new message from a client, adds it to the list, and returns the entire chat history.

Client
Connects to the server via RPC on localhost:1234.

Sends messages to the server by calling the SendMessage method.

Receives and displays the full chat history returned by the server after each message.

🎬 Demo
You can watch a demo video of the project here: [Demo][def]

[def]: https://drive.google.com/drive/folders/1WXcQ2Uk3cwXoUTWRZvFoVoT4zlPX2ZHL?usp=drive_link