package main

import (
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type Client struct {
	id string
	hub *Hub
	socket *websocket.Conn
}

func newClient(hub *Hub, socket *websocket.Conn) *Client {
	return &Client {
		id:     uuid.NewV4().String(),
		hub:    hub,
		socket: socket,
	}
}