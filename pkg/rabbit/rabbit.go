package rabbit

import (
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

const (
	reconnectDelay = 1 * time.Second
)

type RabbitMQ interface {
	NewPublisher(exchange, queue, routingKey string) (*RMQPublisher, error)
	NewConsumer(exchange, queue, routingKey string) (*RMQConsumer, error)
}

type Publisher interface {
	Connect() error
	Publish(message []byte) error
}

type Consumer interface {
	Consume(handler func(message amqp.Delivery))
}

type RMQTransport struct {
	connectionString string
}

func New(connString string) *RMQTransport {
	return &RMQTransport{
		connectionString: connString,
	}
}

func connect(connectionString, exchange, queue, routingKey string) (*amqp.Channel, error) {
	connection, err := amqp.DialConfig(connectionString, amqp.Config{})
	if err != nil {
		return nil, errors.Wrapf(err, "dialing %s failed", connectionString)
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, errors.Wrapf(err, "opening channel failed")
	}

	err = channel.ExchangeDeclare(exchange, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "declare exchange %s failed", exchange)
	}

	_, err = channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "declare queue %s failed", queue)
	}

	err = channel.QueueBind(queue, routingKey, exchange, false, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "bind queue %s to exchange %s (routing key %s) failed", queue, exchange, routingKey)
	}

	return channel, nil
}
