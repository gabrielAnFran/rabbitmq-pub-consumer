package rabbit

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
)

type RabbitModel struct {
}

var Module = fx.Options(fx.Provide(NewRabbit))

func NewRabbit() *RabbitModel {
	return &RabbitModel{}
}

type RabbitInterface interface {
}

func (rabbitModel *RabbitModel) ConsumeMessage(ctx context.Context) error {
	conn, err := amqp.Dial("amqp://fran:cinha@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	mainQueue, err := ch.QueueDeclare(
		"testinho",
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
			log.Printf("Mensagem recebida: %s", string(message.Body))
		}
	}()
	select {}
}
