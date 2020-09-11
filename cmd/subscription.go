package cmd

import (
	"fmt"
	"github.com/felipebool/mockub/internal/client"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

var subscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Short: "handles subscription creation, removal and listing",
}

var createSubscriptionCmd = &cobra.Command{
	Use:   "create",
	Short: "creates subscriptions",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewClient(cmd.Context(), Project, Host, Port)
		if err != nil {
			return err
		}

		t := &client.Topic{ID: TopicID}
		s := &client.Subscription{ID: SubscriptionID}
		if SubscriptionEndpoint != "" {
			s.Endpoint = SubscriptionEndpoint
		}

		if err = c.CreateSubscription(t, s); err != nil {
			return err
		}

		return nil
	},
}

var removeSubscriptionCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes subscriptions",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewClient(cmd.Context(), Project, Host, Port)
		if err != nil {
			return err
		}

		s := &client.Subscription{ID: SubscriptionID}
		if err = c.RemoveSubscription(s); err != nil {
			return err
		}

		return nil
	},
}

var listSubscriptionCmd = &cobra.Command{
	Use:   "list",
	Short: "lists subscriptions",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewClient(cmd.Context(), Project, Host, Port)
		if err != nil {
			return err
		}

		subscriptions, err := c.ListSubscriptions()
		if err != nil {
			return err
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 1, ' ', 0)
		_, _ = fmt.Fprint(w, "ID\tName\tTopic\tEndpoint")
		for _, s := range subscriptions {
			_, _ = fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", s.ID, s.Name, s.TopicID, s.Endpoint)
		}

		_ = w.Flush()
		return nil
	},
}

func init() {
	createSubscriptionCmd.Flags().StringVarP(
		&SubscriptionID,
		"id",
		"i",
		"",
		"id of the subscription to be created",
	)
	createSubscriptionCmd.Flags().StringVarP(
		&TopicID,
		"topic-id",
		"t",
		"",
		"id of the topic to subscribe to",
	)
	createSubscriptionCmd.Flags().StringVarP(
		&SubscriptionEndpoint,
		"endpoint",
		"e",
		"",
		"endpoint to push messages to, when provided subscription would be of type push",
	)

	removeSubscriptionCmd.Flags().StringVarP(
		&SubscriptionID,
		"id",
		"i",
		"",
		"id of the subscription to be removed",
	)

	_ = createSubscriptionCmd.MarkFlagRequired("id")
	_ = removeSubscriptionCmd.MarkFlagRequired("id")
	_ = createSubscriptionCmd.MarkFlagRequired("topic-id")

	subscriptionCmd.AddCommand(createSubscriptionCmd)
	subscriptionCmd.AddCommand(removeSubscriptionCmd)
	subscriptionCmd.AddCommand(listSubscriptionCmd)

	rootCmd.AddCommand(subscriptionCmd)
}
