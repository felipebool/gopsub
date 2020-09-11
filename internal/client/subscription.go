package client

import (
	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
	"time"
)

type Subscription struct {
	ID       string
	Name     string
	TopicID  string
	Endpoint string
}

func (c client) CreateSubscription(topic *Topic, subscription *Subscription) error {
	t := c.pubsubClient.Topic(topic.ID)
	ok, err := t.Exists(c.context)
	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	subConfiguration := pubsub.SubscriptionConfig{
		Topic:            t,
		AckDeadline:      10 * time.Second,
		ExpirationPolicy: 25 * time.Hour,
	}

	if subscription.Endpoint != "" {
		subConfiguration.PushConfig = pubsub.PushConfig{
			Endpoint:             subscription.Endpoint,
		}
	}

	_, err = c.pubsubClient.CreateSubscription(
		c.context,
		subscription.ID,
		subConfiguration,
	)
	if err != nil {
		return err
	}

	return nil
}

func (c client) RemoveSubscription(subscription *Subscription) error {
	s := c.pubsubClient.Subscription(subscription.ID)
	ok, err := s.Exists(c.context)
	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	if err = s.Delete(c.context); err != nil {
		return err
	}

	return nil
}

func (c client) ListSubscriptions() ([]Subscription, error) {
	subscriptions := c.pubsubClient.Subscriptions(c.context)
	list := make([]Subscription, 1)

	for {
		s, err := subscriptions.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		config, err := s.Config(c.context)
		if err != nil {
			return nil, err
		}

		subs := Subscription{
			ID:       s.ID(),
			Name:     s.String(),
			TopicID:  config.Topic.String(),
			Endpoint: config.PushConfig.Endpoint,
		}

		list = append(list, subs)
	}

	return list, nil
}
