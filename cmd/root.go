package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var TopicID string

var Project string

var SubscriptionID string
var SubscriptionEndpoint string

var PublishData string
var PublishDataFromFile string

var Host string
var Port string

var rootCmd = &cobra.Command{
	Use:   "gopsub",
	Short: "an utility to interact with gcloud pubsub emulator",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&Host,
		"host",
		"",
		"localhost",
		"gcloud emulator host",
	)

	rootCmd.PersistentFlags().StringVarP(
		&Port,
		"port",
		"",
		"8085",
		"gcloud emulator port",
	)

	rootCmd.PersistentFlags().StringVarP(
		&Project,
		"project",
		"p",
		"",
		"pubsub project identifier",
	)

	_ = rootCmd.MarkFlagRequired("project")
}
