# Go RPC Chat Project

A simple **client-server chat application** built in Go using **RPC (Remote Procedure Call)**. The project demonstrates a basic chat system where multiple clients can connect to a server, send messages, and view the chat history.

---

## Project Structure

client-server/
├── client/
│ └── client.go
└── server/
└── server.go


- **client/**: Contains the client-side Go program that connects to the server, sends messages, and displays chat history.
- **server/**: Contains the server-side Go program that receives messages from clients, stores them, and returns chat history.

---

## Features

- Connect multiple clients to a single server.
- Send and receive messages in real-time.
- View full chat history after sending a message.
- Thread-safe message handling on the server using mutexes.

---

## Getting Started

### Prerequisites

- Go installed (version 1.18+ recommended)
- Basic understanding of Go programming

---

### Running the Server

1. Navigate to the `server` directory:

```bash
cd server

Run the server:


go run server.go


You should see:

Chat server is running on port 1234...
The server is now ready to accept client connections.



Running the Client

Open a new terminal and navigate to the client directory:

cd client


Run the client:

go run client.go


Enter your name when prompted.

Type messages and press Enter to send.

Type exit to quit the chat

Enter your name: Alice
Type messages to send. Type 'exit' to quit.


> Hello, everyone!
--- Chat History ---
Alice: Hello, everyone!
-------------------

> exit
Exiting...

Demo Video

You can watch a demo video of the project here:
Go [Video](https://drive.google.com/drive/folders/1WXcQ2Uk3cwXoUTWRZvFoVoT4zlPX2ZHL?usp=drive_link)

How It Works

Server:

Listens on TCP port 1234.

Maintains a slice of messages (messages []string).

Uses sync.Mutex to ensure thread-safe updates when multiple clients send messages concurrently.

Exposes a SendMessage method via RPC to accept messages from clients and return the chat history.

Client:

Connects to the server via RPC on localhost:1234.

Sends messages to the server using SendMessage.

Receives and displays the full chat history after each message