package client

import (
	"cloud.google.com/go/pubsub"
	"context"
)

func NewClient(c context.Context, projectID string) (*pubsub.Client, error) {
	return pubsub.NewClient(c, projectID)
}
