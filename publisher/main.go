package main

import (
	"context"

	"publish/rabbit"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Queue   string `json:"queue"`   // Nome da fila
	Message string `json:"message"` // Conteúdo da mensagem
}

func main() {

	// Inicializa o roteador Gin
	router := gin.Default()
	router.POST("/publish", publish)
	gin.SetMode("debug")

	// Inicia o servidor em localhost:8080
	router.Run(":8080")
}

func publish(c *gin.Context) {
	var msg Message

	// Faz o bind dos dados JSON para a struct Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		panic(err)
	}

	// Inicia a conexão com o RabbitMQ e publica a mensagem
	err := rabbit.StartConnect(context.Background(), &msg.Message, &msg.Queue)
	if err != nil {
		panic(err)
	}

	// Responde com mensagem de sucesso
	c.JSON(200, gin.H{"message": "Mensagem enviada"})
}
