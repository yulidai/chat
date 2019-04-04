package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {}

func Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if(err != nil) {
		fmt.Printf("Error occured: %s", err)
		return
	}

	for {
		conn.WriteMessage(websocket.TextMessage, []byte("hahaha"))
		time.Sleep(time.Duration(2) * time.Second)
	}
}