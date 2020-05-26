package rabbit

import (
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// RMQConsumer rabbitmq consumer
type RMQConsumer struct {
	connectionString string
	queue            string
	routingKey       string
	exchange         string
}

func (c *RMQTransport) NewConsumer(exchange, queue, routingKey string) (*RMQConsumer, error) {
	consumer := &RMQConsumer{
		connectionString: c.connectionString,
		queue:            queue,
		routingKey:       routingKey,
		exchange:         exchange,
	}

	return consumer, nil
}

func (c *RMQConsumer) Consume(handler func(message amqp.Delivery)) {
	for {
		messages, err := c.consume()
		if err != nil {
			log.WithError(err).Error("Consumer connect failed")
			time.Sleep(reconnectDelay)
			continue
		}

		log.WithField("exchange", c.exchange).WithField("queue", c.queue).Debug("starting consumer")
		for delivery := range messages {
			handler(delivery)
		}

		log.WithField("exchange", c.exchange).WithField("queue", c.queue).Debug("consumer reconnect ...")
		time.Sleep(reconnectDelay)
	}
}

func (c *RMQConsumer) consume() (<-chan amqp.Delivery, error) {
	ch, err := connect(c.connectionString, c.exchange, c.queue, c.routingKey)
	if err != nil {
		return nil, err
	}

	deliveries, err := ch.Consume(c.queue, "", false, true, false, false, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "can't start listening messages from amqp channel")
	}

	return deliveries, nil
}
