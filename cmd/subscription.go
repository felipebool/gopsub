package cmd

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"google.golang.org/api/iterator"
	"os"
	"time"

	"github.com/spf13/cobra"
	"text/tabwriter"
)

var subscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var createSubscriptionCmd = &cobra.Command{
	Use:   "create",
	Short: "creates a subscription",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetGCloudClient(cmd.Context(), ProjectID)
		
		t := client.Topic(TopicID)
		ok, err := t.Exists(cmd.Context())
		if err != nil {
			fmt.Printf("unable to retrieve topic %s - %v\n", TopicID, err)
			os.Exit(23)
		}

		if !ok {
			fmt.Printf("topic %s does not exist\n", TopicID)
			os.Exit(23)
		}

		subConfiguration := pubsub.SubscriptionConfig{
			Topic:            t,
			AckDeadline:      10 * time.Second,
			ExpirationPolicy: 25 * time.Hour,
		}

		endpoint, err := cmd.Flags().GetString("endpoint")
		if err == nil {
			subConfiguration.PushConfig = pubsub.PushConfig{Endpoint: endpoint}
		}

		_, err = client.CreateSubscription(
			cmd.Context(), 
			SubscriptionID, 
			subConfiguration,
		)

		if err != nil {
			fmt.Printf("unable to create subscription %s to topic %s - %v\n", SubscriptionID, TopicID, err)
			os.Exit(23)
		}

		fmt.Printf("subscription %s to topic %s created\n", SubscriptionID, TopicID)
	},
}

var removeSubscriptionCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes a subscription",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetGCloudClient(cmd.Context(), ProjectID)

		subscription := client.Subscription(SubscriptionID)
		ok, err := subscription.Exists(cmd.Context())
		if err != nil {
			fmt.Println(err)
			os.Exit(23)
		}

		if !ok {
			fmt.Printf("subscription %s does not exist\n", SubscriptionID)
			os.Exit(0)
		}

		if err := subscription.Delete(cmd.Context()); err != nil {
			fmt.Printf("unable to delete subscription %s - %v\n", SubscriptionID, err)
			os.Exit(23)
		}

		fmt.Printf("subscription %s removed\n", SubscriptionID)
	},
}

var listSubscriptionCmd = &cobra.Command{
	Use:   "list",
	Short: "list subscriptions",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetGCloudClient(cmd.Context(), ProjectID)
		subscriptions := client.Subscriptions(cmd.Context())

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 1, ' ', 0)

		_, _ = fmt.Fprintln(w, "ID\tName\tTopic\tEndpoint")
		for {
			s, err := subscriptions.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				fmt.Println(err)
			}

			c, err := s.Config(cmd.Context())
			if err != nil {
				fmt.Println(err)
			}

			topic := c.Topic.String()
			out := fmt.Sprintf("%s\t%s\t%s\t%s\n", s.ID(), s.String(), topic, c.PushConfig.Endpoint)

			_, _ = fmt.Fprint(w, out)
		}
		_ = w.Flush()
	},
}

func init() {
	createSubscriptionCmd.Flags().StringVarP(&SubscriptionID, "id", "i", "", "")
	removeSubscriptionCmd.Flags().StringVarP(&SubscriptionID, "id", "i", "", "")

	createSubscriptionCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "")
	removeSubscriptionCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "")
	listSubscriptionCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "")

	createSubscriptionCmd.Flags().StringVarP(&TopicID, "topic-id", "t", "", "")
	createSubscriptionCmd.Flags().StringVarP(&Endpoint, "endpoint", "e", "", "")
	
	_ = createSubscriptionCmd.MarkFlagRequired("id")
	_ = removeSubscriptionCmd.MarkFlagRequired("id")
	_ = listSubscriptionCmd.MarkFlagRequired("id")
	
	_ = createSubscriptionCmd.MarkFlagRequired("project-id")
	_ = removeSubscriptionCmd.MarkFlagRequired("project-id")
	_ = listSubscriptionCmd.MarkFlagRequired("project-id")
	
	_ = createSubscriptionCmd.MarkFlagRequired("topic-id")

	subscriptionCmd.AddCommand(createSubscriptionCmd)
	subscriptionCmd.AddCommand(removeSubscriptionCmd)
	subscriptionCmd.AddCommand(listSubscriptionCmd)

	rootCmd.AddCommand(subscriptionCmd)
}
