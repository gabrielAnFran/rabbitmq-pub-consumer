package rabbit

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConsumeMessage(ctx context.Context) error {
	conn, err := amqp.Dial("amqp://fran:cinha@rabbitmq/") // Aqui tu coloca o usuário e senha do RabbitMQ que voce setar... né, só pra avisar
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	mainQueue, err := ch.QueueDeclare(
		"testinho", // Esse é o nome da minha fila, tu coloca o nome da tua
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	messageChannel, err := ch.Consume(
		mainQueue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	go func() {
		for message := range messageChannel {
			log.Printf("Mensagem recebida: %s", string(message.Body)) // Aqui tu faz o que tu quiser com a mensagem recebida, no meu caso só printei
		}
	}()
	select {}
}
