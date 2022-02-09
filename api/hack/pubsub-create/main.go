// usage: go run ./main.go -project-id=project-id -topic-id=topic-id
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	start := time.Now()
	fmt.Println("Start..")
	if err := run(); err != nil {
		panic(err)
	}
	fmt.Printf("Done: %s\n", time.Since(start))
}

type app struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func run() error {
	app := &app{}
	projectID := flag.String("project-id", "project-dev", "target project id in gcp")
	topicID := flag.String("topic-id", "", "pubsub topic id to create")
	emulatorPath := flag.String("emulator-path", "", "pubsub emulator path if enable emulator")
	flag.Parse()

	if *emulatorPath != "" {
		os.Setenv("PUBSUB_EMULATOR_HOST", *emulatorPath)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := pubsub.NewClient(ctx, *projectID)
	if err != nil {
		return err
	}
	app.client = client

	fmt.Println("Create Pub/Sub Topic...")
	topic, err := app.createTopic(ctx, *topicID)
	if err != nil {
		return err
	}
	app.topic = topic

	fmt.Println("Create Pub/Sub Subscription...")
	return app.createSubscription(ctx, *topicID)
}

func (a *app) createTopic(ctx context.Context, topicID string) (*pubsub.Topic, error) {
	topic, err := a.client.CreateTopic(ctx, topicID)
	if err == nil {
		return topic, nil
	}
	if status.Code(err) != codes.AlreadyExists {
		return nil, err
	}
	return a.client.Topic(topicID), nil
}

func (a *app) createSubscription(ctx context.Context, subscriptionID string) error {
	conf := pubsub.SubscriptionConfig{Topic: a.topic}
	_, err := a.client.CreateSubscription(ctx, subscriptionID, conf)
	return err
}
