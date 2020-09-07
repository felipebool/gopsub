package cmd

import (
	"cloud.google.com/go/pubsub"
	"github.com/spf13/cobra"
	"time"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		message := pubsub.Message{
			ID:              "123",
			Data:            nil,
			Attributes:      nil,
			PublishTime:     time.Now(),
		}

		client := GetGCloudClient(cmd.Context(), ProjectID)
		t := client.Topic(TopicID)
		t.Publish(cmd.Context(), &message)
	},
}

func init() {
	publishCmd.Flags().StringVarP(&Message, "data", "d", "", "")
	//publishCmd.Flags().StringVarP(&Message, "data-from-file", "f", "", "")

	publishCmd.Flags().StringVarP(&Message, "attributes", "a", "", "")
	publishCmd.Flags().StringVarP(&Message, "project", "p", "", "")
	publishCmd.Flags().StringVarP(&Message, "id", "i", "", "")

	_ = publishCmd.MarkFlagRequired("data")
	_ = publishCmd.MarkFlagRequired("project")
	_ = publishCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(publishCmd)
}
