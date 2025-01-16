package websocket

import (
	"chat-app/internal/domain"
	"chat-app/internal/interactor"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request, chatService *interactor.ChatService) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Ошибка при установке WebSocket-соединения:", err)
		return
	}
	defer conn.Close()
	log.Println("WebSocket соединение установлено")

	client := &domain.Client{
		Send: make(chan []byte),
	}

	chatService.AddClient(client)
	log.Println("Клиент подключен:", client)

	go func() {
		for msg := range client.Send {
			log.Println("Отправка сообщения клиенту:", string(msg))
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Ошибка при отправке сообщения клиенту:", err)
				return
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка при чтении сообщения:", err)
			break
		}

		log.Println("Получено сообщение от клиента:", string(msg))

		message := domain.Message{
			User:    "Пользователь",
			Content: string(msg),
		}
		chatService.Broadcast <- message
		log.Println("Сообщение отправлено всем клиентам:", message)
	}

	chatService.RemoveClient(client)
	log.Println("Клиент отключен:", client)
}
