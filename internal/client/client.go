package client

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
)

type Client interface {
	CreateTopic(topic *Topic) error
	RemoveTopic(topic *Topic) error
	ListTopics() ([]Topic, error)

	CreateSubscription(topic *Topic, subscription *Subscription) error
	RemoveSubscription(subscription *Subscription) error
	ListSubscriptions() ([]Subscription, error)

	Publish(topic *Topic, message *Message) error
}

type client struct {
	projectID    string
	pubsubClient *pubsub.Client
	context      context.Context
}

func NewClient(ctx context.Context, projectID, host, port string) (Client, error) {
	if err := os.Setenv("PUBSUB_EMULATOR_HOST", fmt.Sprintf("%s:%s", host, port)); err != nil {
		return nil, err
	}

	c, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &client{
		projectID:    projectID,
		pubsubClient: c,
		context:      ctx,
	}, nil
}
