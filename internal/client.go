package internal

import (
	"fmt"
	"sync"
)

type Client struct {
	ID   string
	Chan chan string
}

func NewClient(id string, channel chan string) *Client {
	return &Client{
		ID:   id,
		Chan: channel,
	}
}

var clients = make(map[string]*Client)
var clientsLock sync.Mutex

func AddClient(client *Client) {
	clientsLock.Lock()
	defer clientsLock.Unlock()
	clients[client.ID] = client
}

func RemoveClient(clientID string) {
	clientsLock.Lock()
	defer clientsLock.Unlock()
	delete(clients, clientID)
}

func Broadcast(message string) {
	clientsLock.Lock()
	defer clientsLock.Unlock()

	for _, client := range clients {
		select {
		case client.Chan <- message:
		default:
			fmt.Printf("Failed to send message to client %s\n", client.ID)
		}
	}
}
