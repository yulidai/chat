package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Start now!")
	http.HandleFunc("/ws", Handler)
	http.ListenAndServe("localhost:9191", nil)
}