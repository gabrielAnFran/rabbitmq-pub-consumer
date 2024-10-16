package rabbit

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func StartConnect(ctx context.Context, message, queue *string) error {
	connection, err := amqp.Dial("amqp://fran:cinha@rabbitmq:5672") // Aqui tu coloca o usuário e senha do RabbitMQ que voce setar... né, só pra avisar
	if err != nil {
		fmt.Println(err.Error())
	}
	defer connection.Close()
	ch, err := connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	mainQueue, err := ch.QueueDeclare(
		*queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(
		ctx,
		"",
		mainQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(*message),
		},
	)

	if err != nil {
		return err
	}
	return nil
}
