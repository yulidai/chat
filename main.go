package main

import (
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()
	
	http.HandleFunc("/ws", hub.handler)
	http.ListenAndServe("localhost:9191", nil)
}