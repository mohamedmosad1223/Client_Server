تفضل، هذا هو الكود الكامل للملف، جاهز للنسخ واللصق مباشرة:

```markdown
# Go RPC Chat Project

A simple client-server chat application built in Go using RPC (Remote Procedure Call). This project demonstrates a basic chat system where multiple clients can connect to a server, send messages, and view the chat history.

## 🚀 Features

* Connect multiple clients to a single server.
* Send and receive messages in real-time.
* View the full chat history after each message is sent.
* Thread-safe message handling on the server using `sync.Mutex`.

## 📁 Project Structure

```

client-server/
├── client/
│   └── client.go
└── server/
└── server.go

````

* **client/**: Contains the client-side Go program that connects to the server.
* **server/**: Contains the server-side Go program that manages the chat.

## 🏁 Getting Started

### Prerequisites

* [Go](https://golang.org/doc/install) (version 1.18+ recommended)
* A basic understanding of Go programming.

### 🛠️ Installation & Running

1.  **Clone the repository** (أو قم بتنزيل الملفات).

2.  **Run the Server**
    افتح نافذة Terminal وانتقل إلى مجلد `server`:
    ```bash
    cd server
    go run server.go
    ```
    سترى الرسالة التالية، مما يعني أن الخادم جاهز:
    ```
    Chat server is running on port 1234...
    ```

3.  **Run the Client**
    افتح نافذة Terminal **جديدة** وانتقل إلى مجلد `client`:
    ```bash
    cd client
    go run client.go
    ```

4.  **Start Chatting!**
    * سيطلب منك إدخال اسمك.
    * اكتب رسائلك واضغط `Enter` لإرسالها.
    * اكتب `exit` للخروج من المحادثة.

**Example Usage:**
````

Enter your name: Alice
Type messages to send. Type 'exit' to quit.

> Hello, everyone\!
> \--- Chat History ---
> Alice: Hello, everyone\!

-----

> exit
> Exiting...

```

## 🔧 How It Works

### Server
* Listens on TCP port `1234`.
* Maintains a slice of all messages (`messages []string`).
* Uses `sync.Mutex` to ensure thread-safety when multiple clients send messages concurrently.
* Exposes a `SendMessage` method via RPC. This method accepts a new message from a client, adds it to the list, and returns the entire chat history.

### Client
* Connects to the server via RPC on `localhost:1234`.
* Sends messages to the server by calling the `SendMessage` method.
* Receives and displays the full chat history returned by the server after each message.

## 🎬 Demo

You can watch a demo video of the project here:
[Watch Demo](https://drive.google.com/drive/folders/1WXcQ2Uk3cwXoUTWRZvFoVoT4zlPX2ZHL?usp=drive_link)

## 🛠️ Built With

* ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

```