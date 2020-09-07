# Gopsub

This is a simple, command line client for GCloud PubSub emulator. The idea is to help
debugging applications which use GCloud PubSub by providing a simple interface for
topic/subscriptions creation and message publishing.

# Installation
Not done yet.

# Commands
`gopsub` allows you to create, remove and list topics and subscriptions,
and message publishing.

## Topic

### List topics
` PUBSUB_EMULATOR_HOST=localhost:8085 go run main.go topic list` shows help for
topic listing.

#### Required flags
* `--project-id` or `-p`: project id

### Create topic
` PUBSUB_EMULATOR_HOST=localhost:8085 go run main.go topic create` shows help for
topic creation.

#### Required flags
* `--project-id` or `-p`: project id
* `--id` or `-i`: new topic id

### Remove topic
` PUBSUB_EMULATOR_HOST=localhost:8085 go run main.go topic remove` shows help for
topic removal.

#### Required flags
* `--project-id` or `-p`: project id
* `--id` or `-i`: topic id to be removed


## Subscription

### List subscriptions
` PUBSUB_EMULATOR_HOST=localhost:8085 go run main.go subscription list` shows help for
subscription listing.

#### Required flags
* `--project-id` or `-p`: project id

### Create subscription
` PUBSUB_EMULATOR_HOST=localhost:8085 go run main.go subscription create` shows help for
subscription creation.

#### Required flags
* `--project-id` or `-p`: project id
* `--id` or `-i`: new subscription id
* `--topic-id` or `-t`: topic id to create subscription to

#### Optional fields
* `--endpoint` or  `-e`: when present, subscription will be of type push and push
messages to endpoint

### Remove topic
` PUBSUB_EMULATOR_HOST=localhost:8085 go run main.go topic remove` shows help for
subscription removal.

#### Required fields
* `--project-id` or `-p`: project id
* `--id` or `-i`: subscription id to be removed

## Publish
Coming