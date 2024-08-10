package notification

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

type LoginServiceManager struct {
	// Events are pushed to this channel by the main events-gathering routine
	message chan string

	// New client connections
	newClients chan chan string

	// Closed client connections
	closedClients chan chan string

	// Total client connections
	clients map[chan string]bool

	once sync.Once
}

var ourloginServicesManger = new(LoginServiceManager)

func GetLoginServiceManager() *LoginServiceManager {
	ourloginServicesManger.once.Do(func() {
		ourloginServicesManger.message = make(chan string)
		ourloginServicesManger.newClients = make(chan chan string)
		ourloginServicesManger.closedClients = make(chan chan string)
		ourloginServicesManger.clients = make(map[chan string]bool)

		go ourloginServicesManger.listen()
	})
	return ourloginServicesManger
}

func (manager *LoginServiceManager) RegisterClient(client chan string) {
	manager.newClients <- client
}

func (manager *LoginServiceManager) UnRegisterClient(client chan string) {
	manager.closedClients <- client
}

func (manager *LoginServiceManager) SendMessage(msg any) {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		jsonData, _ = json.Marshal(gin.H{"Error": err})
	}
	manager.message <- string(jsonData)
}

func (manager *LoginServiceManager) listen() {
	for {
		select {
		// Add new available client
		case client := <-manager.newClients:
			manager.clients[client] = true
			log.Printf("Client loginServicesEvent added. %d registered clients", len(manager.clients))
		// Remove closed client
		case client := <-manager.closedClients:
			delete(manager.clients, client)
			close(client)
			log.Printf("Removed loginServicesEvent client. %d registered clients", len(manager.clients))
		// Broadcast message to client
		case eventMsg := <-manager.message:
			for clientMessageChan := range manager.clients {
				clientMessageChan <- eventMsg
			}
		}
	}
}
