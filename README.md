# RabbitMQ Project

This project contains a RabbitMQ publisher, a consumer, and a docker-compose file to set up a RabbitMQ service.

## Running the Project

To run the entire project, including the publisher, consumer, and RabbitMQ service, follow these steps:

1. Make sure Docker is installed.
2. Navigate to the project directory.
3. Run `docker compose up`.

This will start all services as defined in the `docker-compose.yml` file.

## Publisher

The publisher is configured in the `docker-compose.yml` to run in a Docker container named `publisher_service`, exposed on port 8080.

## Consumer

The consumer is configured in the `docker-compose.yml` to run in a Docker container named `consumer_service`, exposed on port 8082.

## RabbitMQ Instance

The `docker-compose.yml` file defines the configuration of a RabbitMQ service in a Docker environment using the `rabbitmq:management` image. It sets up a container named `test_rabbitmq` with ports 5672 and 15672 exposed, and sets the environment variables `RABBITMQ_DEFAULT_USER` to `fran` and `RABBITMQ_DEFAULT_PASS` to `cinha`.

## Networks

The services are configured to use a Docker network called `my_network`, which uses the `bridge` driver.

## Testing

To test the functionality, use the following curl http action, which can be found in `publisher/test.http`:

```bash
POST http://localhost:8080/publish
Content-Type: application/json

{
    "queue": "testinho",
    "message": "Good good not bad bad!"
}

```