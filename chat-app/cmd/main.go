package main

import (
	"chat-app/internal/interactor"
	"chat-app/internal/interfaces/http"
)

func main() {
	chatService := interactor.NewChatService()
	http.StartServer(chatService)
}
