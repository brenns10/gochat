/*
Package gochat implements a simple TCP-based chat protocol.

This chat protocol is very simple. The client and server exchange commands,
which are JSON objects.
*/
package gochat

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

/*
ChatServer holds the clients of the chat server.
*/
type ChatServer struct {
	clients []net.Conn
}

/*
Create a new chat server!
*/
func NewChatServer() ChatServer {
	return ChatServer{clients: []net.Conn{}}
}

/*
Run the chat server
*/
func (cs *ChatServer) Run(addr string) bool {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Encountered error accepting connection: ", err)
			return false
		}
		cs.clients = append(cs.clients, conn)
		go cs.runClient(conn)
	}
}

/*
Broadcast the given message to all connected clients.
*/
func (cs *ChatServer) broadcastToClients(msg map[string]string) {
	bytes, _ := json.Marshal(msg)

	for _, client := range cs.clients {
		client.Write(bytes)
		client.Write([]byte{'\n'})
	}
}

/*
Run a goroutine for a client.
*/
func (cs *ChatServer) runClient(client net.Conn) {
	decoder := json.NewDecoder(client)
	for decoder.More() {
		m := make(map[string]string)
		decoder.Decode(&m)
		fmt.Println("Decoded message", m)
		if m["cmd"] == "msg" {
			cs.broadcastToClients(m)
		}
	}
}

/*
Represents the data for a Client on the server.
*/
type ChatClient struct {
	conn net.Conn
}

/*
Create a new chat client.
*/
func NewChatClient(addr string) *ChatClient {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("error encountered connecting")
	}
	return &ChatClient{conn: conn}
}

/*
Routine that reads messages from the server and prints them out.
*/
func (cc *ChatClient) runRead() {
	decoder := json.NewDecoder(cc.conn)
	for decoder.More() {
		m := make(map[string]string)
		decoder.Decode(&m)
		if m["cmd"] == "msg" {
			fmt.Println(m["msg"])
		}
	}
}

/*
The main chat client function.  Loops reading input and sending it to server.
*/
func (cc *ChatClient) Run() {
	go cc.runRead()
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		msg := make(map[string]string)
		msg["cmd"] = "msg"
		msg["msg"] = str
		bytes, _ := json.Marshal(msg)
		cc.conn.Write(bytes)
		cc.conn.Write([]byte{'\n'})
	}
}
