package client

import "google.golang.org/api/iterator"

type Topic struct {
	ID string
	Name string
}

func (c client) CreateTopic(topic *Topic) error {
	t := c.pubsubClient.Topic(topic.ID)
	ok, err := t.Exists(c.context)
	if err != nil {
		return err
	}

	if ok {
		return nil
	}

	_, err = c.pubsubClient.CreateTopic(c.context, topic.ID)
	if err != nil {
		return err
	}

	return nil
}

func (c client) RemoveTopic(topic *Topic) error {
	t := c.pubsubClient.Topic(topic.ID)
	ok, err := t.Exists(c.context)
	if err != nil {
		return err
	}

	if !ok {
		return nil
	}

	return t.Delete(c.context)
}

func (c client) ListTopics() ([]Topic, error) {
	topics := c.pubsubClient.Topics(c.context)
	list := make([]Topic, 1)

	for {
		t, err := topics.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		list = append(list, Topic{ID: t.ID(), Name: t.String()})
	}

	return list, nil
}
