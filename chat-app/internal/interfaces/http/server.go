package http

import (
	"chat-app/internal/interactor"
	"chat-app/internal/interfaces/websocket"
	"fmt"
	"log"
	"net/http"
)

func StartServer(chatService *interactor.ChatService) {

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.HandleWebSocket(w, r, chatService)
	})

	fmt.Println("Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
