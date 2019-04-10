package main
import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Hub
type Hub struct {
	clientList []*Client
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub {
		clientList: make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if(err != nil) {
		fmt.Printf("Error occured: %s", err)
		return
	}

	hub.register <- newClient(hub, conn)
}

func (hub *Hub) broadcast(msg string) {
	for _, client := range hub.clientList {
		client.socket.WriteMessage(websocket.TextMessage, []byte("Broadcast: " + msg))
	}
}

func (hub *Hub) onConnect(client *Client) {
	fmt.Println("On new connect")
	hub.broadcast("new client is coming in")
	hub.clientList = append(hub.clientList, client)
}

func (hub *Hub) onDisconnect(clientLeave *Client) {
	index := -1
	for i, client := range hub.clientList {
		if(client.id != clientLeave.id) {
			continue
		}

		index = i
		break
	}

	copy(hub.clientList[index:], hub.clientList[index+1:])
	hub.clientList[len(hub.clientList) - 1] = nil
	hub.clientList = hub.clientList[:len(hub.clientList) - 1]

	hub.broadcast("old client is leave")
}

func (hub *Hub) onMessage(msg []byte) {
	hub.broadcast(string(msg))
}

func (hub *Hub) run() {
	for {
		select {
		case client := <- hub.register:
			hub.onConnect(client)
		case client := <- hub.unregister:
			hub.onDisconnect(client)
		}
	}
}