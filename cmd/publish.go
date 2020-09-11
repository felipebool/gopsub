package cmd

import (
	"errors"
	"github.com/felipebool/mockub/internal/client"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "handles message publishing",
	RunE: func(cmd *cobra.Command, args []string) error {
		dataToSend, err := getDataForPublish(PublishData, PublishDataFromFile)
		if err != nil {
			return err
		}

		t := &client.Topic{ID: TopicID}
		m := &client.Message{Content: dataToSend}

		c, err := client.NewClient(cmd.Context(), Project, Host, Port)
		if err != nil {
			return err
		}

		return c.Publish(t, m)
	},
}

func getDataForPublish(message, messageFilePath string) ([]byte, error) {
	if message == "" && messageFilePath == "" {
		return nil, errors.New("you have to provide at least one source of data")
	}

	if message != "" {
		return []byte(message), nil
	}

	d, err := ioutil.ReadFile(messageFilePath)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func init() {
	publishCmd.Flags().StringVarP(
		&PublishData,
		"data",
		"d",
		"",
		"data to be published (has precedence over --data-from-file)",
	)

	publishCmd.Flags().StringVarP(
		&TopicID,
		"topic-id",
		"t",
		"",
		"id of topic to publish messages to",
	)

	publishCmd.Flags().StringVarP(
		&PublishDataFromFile,
		"data-from-file",
		"",
		"",
		"path to a file to read the data to publish",
	)

	_ = publishCmd.MarkFlagRequired("topic-id")

	rootCmd.AddCommand(publishCmd)
}
