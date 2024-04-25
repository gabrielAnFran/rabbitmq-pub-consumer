package rabbit

import (
	"context"

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

func (rabbitModel *RabbitModel) StartConnect(ctx context.Context) error {
	connection, err := amqp.Dial("amqp://fran:cinha@localhost:5672")
	if err != nil {
		return err
	}
	defer connection.Close()
	ch, err := connection.Channel()
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

	err = ch.PublishWithContext(
		ctx,
		"",
		mainQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("batata"),
		},
	)

	if err != nil {
		return err
	}
	return nil
}
