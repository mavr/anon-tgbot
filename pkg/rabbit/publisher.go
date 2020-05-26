package rabbit

import (
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RMQPublisher struct {
	connectionString string
	channel          *amqp.Channel
	closeCh          chan *amqp.Error

	queue      string
	routingKey string
	exchange   string
}

// NewPublisher create new publisher
func (c *RMQTransport) NewPublisher(exchange, queue, routingKey string) (*RMQPublisher, error) {
	publisher := &RMQPublisher{
		connectionString: c.connectionString,
		exchange:         exchange,
		routingKey:       routingKey,
		queue:            queue,
	}

	return publisher, nil
}

// Connect create connect for publisher
func (p *RMQPublisher) Connect() error {
	if err := p.connect(); err != nil {
		return errors.Wrapf(err, "connect to amqp failed")
	}

	go p.keepConnection()

	return nil
}

func (p *RMQPublisher) connect() error {
	ch, err := connect(p.connectionString, p.exchange, p.queue, p.routingKey)
	if err != nil {
		return errors.Wrapf(err, "can't open connection")
	}

	p.closeCh = make(chan *amqp.Error)
	ch.NotifyClose(p.closeCh)
	p.channel = ch

	return nil
}

func (p *RMQPublisher) keepConnection() {
	go func() {
		for {
			closeErr := <-p.closeCh
			logrus.WithError(closeErr).Error("[AMQP] Producer lost connection. Trying to reconnect...")
			for {
				err := p.connect()
				if err != nil {
					logrus.WithError(err).Error("[AMQP] Producer can't connect")
					time.Sleep(reconnectDelay)
					continue
				}

				break
			}

			logrus.Info("[AMQP] Producer reconnected")
		}
	}()
}

func (p *RMQPublisher) Publish(message []byte) error {
	err := p.channel.Publish(
		p.exchange,
		p.routingKey,
		true,
		false,
		amqp.Publishing{
			Body:         message,
			DeliveryMode: amqp.Persistent,
			Priority:     1,
		},
	)

	if err != nil {
		return errors.Wrapf(err, "can't publish message")
	}

	return nil
}
