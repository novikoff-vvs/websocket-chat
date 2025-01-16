package interactor

import (
	"chat-app/internal/domain"
	"log"
)

type ChatService struct {
	Clients   map[*domain.Client]bool
	Broadcast chan domain.Message
}

func NewChatService() *ChatService {
	cs := &ChatService{
		Clients:   make(map[*domain.Client]bool),
		Broadcast: make(chan domain.Message),
	}
	go cs.BroadcastMessages()
	return cs
}

func (cs *ChatService) AddClient(client *domain.Client) {
	cs.Clients[client] = true
	log.Println("Клиент добавлен:", client)
}

func (cs *ChatService) RemoveClient(client *domain.Client) {
	delete(cs.Clients, client)
	log.Println("Клиент удален:", client)
}

func (cs *ChatService) BroadcastMessages() {
	for message := range cs.Broadcast {
		log.Println("Отправка сообщения всем клиентам:", message)
		for client := range cs.Clients {
			select {
			case client.Send <- []byte(message.Content):
			default:
				log.Println("Ошибка при отправке сообщения клиенту:", client)
				delete(cs.Clients, client)
			}
		}
	}
}
