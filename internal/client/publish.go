package client

import (
	"cloud.google.com/go/pubsub"
	"log"
)

type Message struct {
	Content []byte
}

func (c client) Publish(topic *Topic, message *Message) error {
	t := c.pubsubClient.Topic(topic.ID)

	ok, err := t.Exists(c.context)
	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	res := t.Publish(c.context, &pubsub.Message{Data: message.Content})

	_, err = res.Get(c.context)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
