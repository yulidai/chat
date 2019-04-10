package main

import (
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	id string
	socket *websocket.Conn
}

func newClient(socket *websocket.Conn) *Client {
	return &Client {
		id:     uuid.NewV4().String(),
		socket: socket,
	}
}