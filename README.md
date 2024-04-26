# Projeto RabbitMQ

Este projeto contém um publisher RabbitMQ, um consumer e um docker-compose para levantar um serviço do RabbitMQ.

## Publisher

Para rodar o publisher que levanta um servidor que recebe a mensagem em localhost:8080/publish, siga os passos abaixo:

1. Navegue até o diretório do publisher.
2. `go mod tidy`
3. `go run main.go`

## Consumer

O consumer consome da fila "testinho". Para executar o consumer, siga as instruções:

1. Vá para o diretório do consumer.
2. `go mod tidy`
3. `go run main.go`

## Instância do RabbitMQ

O arquivo docker-compose.yml define a configuração de um serviço RabbitMQ em um ambiente Docker usando a imagem rabbitmq:management. Ele configura um contêiner chamado test_rabbitmq com portas expostas 5672 e 15672, e define as variáveis de ambiente RABBITMQ_DEFAULT_USER como fran e RABBITMQ_DEFAULT_PASS como cinha.

1. Certifique-se de ter o Docker instalado.
2. Navegue até o diretório do projeto.
3. `docker compose up` 

## Teste

Para testar o funcionamento, utilize o seguinte comando curl:


```bash
curl --location 'http://localhost:8080/publish' \
--header 'Content-Type: application/json' \
--data '
{
"queue": "testinho",
"message": "{\"queue\":\"testinho\",\"message\": \"test\"}"
}
'
```