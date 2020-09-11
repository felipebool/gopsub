# Gopsub

This is a simple, command line client for GCloud PubSub emulator. The idea is to help
local debug of applications using GCloud PubSub, by providing a simple interface for
topic/subscription creation/removal and message publishing.

# Installation
Run `go get github.com/felipebool/gopsub`. 

# Commands
`gopsub` allows you to create, remove and list topics and subscriptions,
and publish messages.

## Global required flags
* `--project` or `-p`: project id

## Global optional flags
* `--host`: gcloud pubsub emulator host, default value is `localhost`
* `--post`: gcloud pubsub emulator port, default value is `8085`

## Topic

### List topics
The following command lists all topics for project `my-project`

`gopsub --project my-project topic list`

### Create topic
The following command creates a topic called `topic1` in project `my-project`

`gopsub --project my-project topic create --id topic1`

### Remove topic
The following command removes topic `topic1` from project `my-project`

`gopsub --project my-project topic remove --id topic1`

## Subscription

### List subscriptions
The following command lists all subscriptions for project `my-project`

`gopsub --project my-project subscription list`

### Create subscription
The following command creates a subscription called `subscription1` in project `my-project`

`gopsub --project my-project subscription create --id subscription1`

To create a push subscription, just provide a valid endpoint using `--endpoint` or `-e`

`gopsub --project my-project subscription create --id subscription1 --endpoint http://localhost/push`

### Remove subscription
The following command removes subscription `subscription1` from project `my-project`

`gopsub --project my-project subscription remove --id subscription1`

## Publish
The following command publishes a message passed by `--data` to topic `topic1`

`gopsub --project my-project publish --topic-id topic1 --data '{"username":"xyz","password":"xyz"}'`

If you want to publish more complex data, you can use `--data-from-file` passing
a path to a file. `gopsub` will read the content and publish it.

`gopsub --project my-project publish --topic-id topic1 --data-from-file /tmp/lol.json`

if `--data` and `--data-from-file` are both provided, the value passed through `--data`
will be used.