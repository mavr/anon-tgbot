package delivery

import (
	"encoding/json"

	"github.com/mavr/anonymous-mail/models"
	"github.com/mavr/anonymous-mail/pkg/rabbit"
	"github.com/streadway/amqp"
)

// ChatAmqpDelivery amqp consumer and producer for chat
type ChatAmqpDelivery struct {
	newChatPub  rabbit.Publisher
	newChatCons rabbit.Consumer
}

// ChatAMQPConfig configurator for rabbitmq queues
type ChatAMQPConfig struct {
	NewChatExchange string
	NewChatRK       string
	NewChatQueue    string
}

// NewChatAmqpDelivery create new NewChatAmqpDelivery object
func NewChatAmqpDelivery(mq rabbit.RabbitMQ, c ChatAMQPConfig) (*ChatAmqpDelivery, error) {
	NewChatPubl, err := mq.NewPublisher(c.NewChatExchange, c.NewChatQueue, c.NewChatRK)
	if err != nil {
		return nil, err
	}

	NewChatCons, err := mq.NewConsumer(c.NewChatExchange, c.NewChatQueue, c.NewChatRK)
	if err != nil {
		return nil, err
	}

	return &ChatAmqpDelivery{
		newChatPub:  NewChatPubl,
		newChatCons: NewChatCons,
	}, nil
}

// RegisterPublisher create connection for rabbit amqp publisher.
func (q *ChatAmqpDelivery) RegisterPublisher() error {
	return q.newChatPub.Connect()
}

// NewChatNotificate need in case when new user connect to bot.
// This method publish event about this in queue.
func (q *ChatAmqpDelivery) NewChatNotificate(c *models.Chat) error {
	b, _ := json.Marshal(c)
	return q.newChatPub.Publish(b)
}

// NewChatSubscribe subcriber
func (q *ChatAmqpDelivery) NewChatSubscribe(handle func(c *models.Chat) error) {
	convertHandle := func(m amqp.Delivery) {
		c := models.Chat{}
		_ = json.Unmarshal(m.Body, &c)

		if err := handle(&c); err != nil {
			_ = m.Ack(false)

			return
		}

		_ = m.Ack(true)
	}

	q.newChatCons.Consume(convertHandle)
}
