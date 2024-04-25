package main

import (
	"context"

	"github.com/gabrielAnFran/rabbitmq-pub-consumer/publisher/rabbit"

	"go.uber.org/fx"
)

func Register(lifeCycle fx.Lifecycle, rabbit *rabbit.RabbitModel) { // Define a função Register que recebe fx.Lifecycle e RabbitModel como parâmetros
	lifeCycle.Append(fx.Hook{ // Adiciona um hook ao ciclo de vida
		OnStart: func(ctx context.Context) error { // Define uma função anônima a ser executada no início
			return rabbit.StartConnect(ctx) // Chama o método StartConnect no RabbitModel
		},
		OnStop: nil, 
	})
}

//Isso cria uma nova aplicação FX (framework de injeção de dependência) que 
// inclui o módulo do RabbitMQ e invoca a função Register. Em seguida, 
// inicia a aplicação em segundo plano.
func main() {
	app := fx.New( 
		rabbit.Module,      
		fx.Invoke(Register),
	)
	app.Start(context.Background()) 
}
