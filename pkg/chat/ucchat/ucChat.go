package ucchat

import (
	"errors"

	"github.com/mavr/anonymous-mail/models"
	"github.com/mavr/anonymous-mail/pkg/chat"
	"github.com/mavr/anonymous-mail/pkg/rabbit"

	"github.com/mavr/anonymous-mail/pkg/chat/delivery"
)

const (
	newChatQueueName  = "new_chat"
	newChatRoutingKey = ""

	chNChatSize = 64
)

// Usecase implements chat Usecase interface
type Usecase struct {
	delivery *delivery.ChatAmqpDelivery
	store    chat.Repository

	chNChat chan *models.Chat
}

// New creates new chat Usecase
func New(mq rabbit.RabbitMQ, store chat.Repository, exchange string) (*Usecase, error) {
	delivery, err := delivery.NewChatAmqpDelivery(mq, delivery.ChatAMQPConfig{
		NewChatExchange: exchange,
		NewChatQueue:    newChatQueueName,
		NewChatRK:       newChatRoutingKey,
	})
	if err != nil {
		return nil, err
	}

	return &Usecase{
		delivery: delivery,
		store:    store,
	}, nil
}

// NewChatNotificate send in amqp queue info about new user's chat
func (uc *Usecase) NewChatNotificate(c *models.Chat) error {
	return uc.delivery.NewChatNotificate(c)
}

// GetChatByUserUID return chat model by user uid
func (uc *Usecase) GetChatByUserUID(userUID string) (*models.Chat, error) {
	return uc.store.GetChat(userUID)
}

// SaveNewChat store new chat.
func (uc *Usecase) SaveNewChat(c *models.Chat) error {
	return uc.store.SetChat(c)
}

// RegisterPublisher create connect for new chat notification.
func (uc *Usecase) RegisterPublisher() error {
	return uc.delivery.RegisterPublisher()
}

// RegisterSubscriber create connect for consumer
// and start receive messages in inside channel.
func (uc *Usecase) RegisterSubscriber() error {
	if uc.chNChat == nil {
		uc.chNChat = make(chan *models.Chat, chNChatSize)
	}

	handler := func(c *models.Chat) error {
		uc.chNChat <- c

		return nil
	}

	go uc.delivery.NewChatSubscribe(handler)

	return nil
}

// GetNewChatNotificate this is blocking operation which return
// notification about new chat.
func (uc *Usecase) GetNewChatNotificate() (*models.Chat, error) {
	if uc.chNChat == nil {
		return nil, errors.New("subscriber not initialized")
	}

	c, ok := <-uc.chNChat
	if !ok && c == nil {
		return nil, errors.New("subscriber down")
	}

	return c, nil
}

// GetNewChatNotificateChan this is blocking operation which return
// notification about new chat.
func (uc *Usecase) GetNewChatNotificateChan() (<-chan *models.Chat, error) {
	if uc.chNChat == nil {
		return nil, errors.New("subscriber not initialized")
	}

	return uc.chNChat, nil
}
