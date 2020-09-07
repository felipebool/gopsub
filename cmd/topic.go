package cmd

import (
	"fmt"
	"google.golang.org/api/iterator"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var topicCmd = &cobra.Command{
	Use:   "topic",
	Short: "Manages topics",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var createTopicCmd = &cobra.Command{
	Use:   "create",
	Short: "creates topic",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetGCloudClient(cmd.Context(), ProjectID)

		t := client.Topic(TopicID)
		ok, err := t.Exists(cmd.Context())
		if err != nil {
			fmt.Println(err)
		}

		if ok {
			fmt.Printf("topic %s already exists\n", TopicID)
			os.Exit(0)
		}

		_, err = client.CreateTopic(cmd.Context(), TopicID)
		if err != nil {
			fmt.Printf("unable to create topic %s - %v\n", TopicID, err)
			os.Exit(23)
		}

		fmt.Printf("topic %s create\n", TopicID)
	},
}

var removeTopicCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove topic",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetGCloudClient(cmd.Context(), ProjectID)

		t := client.Topic(TopicID)
		ok, err := t.Exists(cmd.Context())
		if err != nil {
			fmt.Println(err)
			os.Exit(23)
		}

		if !ok {
			fmt.Printf("topic %s does not exist\n", TopicID)
			os.Exit(0)
		}

		if err := t.Delete(cmd.Context()); err != nil {
			fmt.Printf("unable to delete topic %s - %v\n", TopicID, err)
			os.Exit(23)
		}

		fmt.Printf("topic %s removed\n", TopicID)
	},
}

var listTopicCmd = &cobra.Command{
	Use:   "list",
	Short: "list topics",
	Run: func(cmd *cobra.Command, args []string) {
		client := GetGCloudClient(cmd.Context(), ProjectID)

		topics := client.Topics(cmd.Context())

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 0, 1, ' ', 0)

		_, _ = fmt.Fprintln(w, "ID\tName")
		for {
			t, err := topics.Next()
			if err == iterator.Done {
				break
			}

			if err != nil {
				fmt.Println(err)
			}

			out := fmt.Sprintf("%s\t%s\n", t.ID(), t.String())

			_, _ = fmt.Fprint(w, out)
		}

		_ = w.Flush()
	},
}

func init() {
	createTopicCmd.Flags().StringVarP(&TopicID, "id", "i", "", "")
	removeTopicCmd.Flags().StringVarP(&TopicID, "id", "i", "", "")

	createTopicCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "")
	removeTopicCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "")
	listTopicCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "")

	_ = createTopicCmd.MarkFlagRequired("id")
	_ = removeTopicCmd.MarkFlagRequired("id")
	_ = createTopicCmd.MarkFlagRequired("project-id")
	_ = removeTopicCmd.MarkFlagRequired("project-id")
	_ = listTopicCmd.MarkFlagRequired("project-id")

	topicCmd.AddCommand(createTopicCmd)
	topicCmd.AddCommand(removeTopicCmd)
	topicCmd.AddCommand(listTopicCmd)

	rootCmd.AddCommand(topicCmd)
}
