package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool { return true },
}
var clientList = make([]*Client, 0)

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if(err != nil) {
		fmt.Printf("Error occured: %s", err)
		return
	}

	// 通知当前 socket
	for _, c := range clientList {
		c.socket.WriteMessage(websocket.TextMessage, []byte("new member is coming"))
	}

	// 添加新 socket
	clientList = append(clientList, newClient(conn))
}