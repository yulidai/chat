package main

import (
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"fmt"
)

type Client struct {
	id string
	hub *Hub
	socket *websocket.Conn
}

func newClient(hub *Hub, socket *websocket.Conn) *Client {
	client := &Client {
		id:     uuid.NewV4().String(),
		hub:    hub,
		socket: socket,
	}

	client.run()

	return client
}

func (client *Client) read() {
	for {
		_, msg, err := client.socket.ReadMessage()
		if err != nil {
			fmt.Printf("error occured when read message from client: %s", err)
			break
		}
		client.hub.onMessage(msg, client)

		fmt.Printf("msg: %s\n", msg)
	}
}

func (client *Client) run() {
	go client.read()
}