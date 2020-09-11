package cmd

import (
	"fmt"
	"github.com/felipebool/mockub/internal/client"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "handles topic creation, removal and listing",
}

var createTopicCmd = &cobra.Command{
	Use:   "create",
	Short: "creates topics",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewClient(cmd.Context(), Project, Host, Port)
		if err != nil {
			return err
		}

		if err = c.CreateTopic(&client.Topic{ID: TopicID}); err != nil {
			return err
		}

		return nil
	},
}

var removeTopicCmd = &cobra.Command{
	Use:   "remove",
	Short: "removes topic",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewClient(cmd.Context(), Project, Host, Port)
		if err != nil {
			return err
		}

		if err = c.RemoveTopic(&client.Topic{ID: TopicID}); err != nil {
			return err
		}

		return nil
	},
}

var listTopicCmd = &cobra.Command{
	Use:   "list",
	Short: "lists topics",
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewClient(cmd.Context(), Project, Host, Port)
		if err != nil {
			return err
		}

		topics, err := c.ListTopics()
		if err != nil {
			return err
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 1, ' ', 0)
		_, _ = fmt.Fprint(w, "ID\tName")
		for _, t := range topics {
			_, _ = fmt.Fprintf(w, "%s\t%s\n", t.ID, t.Name)
		}

		_ = w.Flush()
		return nil
	},
}

func init() {
	createTopicCmd.Flags().StringVarP(
		&TopicID,
		"id",
		"i",
		"",
		"id of the topic to be created",
	)

	removeTopicCmd.Flags().StringVarP(
		&TopicID,
		"id",
		"i",
		"",
		"id of the topic to be removed",
	)

	_ = createTopicCmd.MarkFlagRequired("id")
	_ = removeTopicCmd.MarkFlagRequired("id")

	topicCmd.AddCommand(createTopicCmd)
	topicCmd.AddCommand(removeTopicCmd)
	topicCmd.AddCommand(listTopicCmd)

	rootCmd.AddCommand(topicCmd)
}
