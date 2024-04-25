package main

import (
	"consume/rabbit"
	"context"
	"log"

	"go.uber.org/fx"
)

func Register(lifeCycle fx.Lifecycle, rabbitModel *rabbit.RabbitModel) {
	lifeCycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := rabbitModel.ConsumeMessage(ctx)
				if err != nil {
					log.Println(err)
				}
			}()
			return nil
		},
		OnStop: nil,
	})
}

func main() {
	app := fx.New(
		rabbit.Module,
		fx.Invoke(Register),
	)
	app.Run()
}
