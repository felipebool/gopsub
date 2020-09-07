package cmd

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"github.com/felipebool/mockub/internal/client"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var TopicID string
var SubscriptionID string
var ProjectID string
var Endpoint string
var Message string

var rootCmd = &cobra.Command{
	Use:   "mocksub",
	Short: "An utility to interact with gcloud pubsub emulator",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mocksub.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mocksub" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mocksub")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func GetGCloudClient(ctx context.Context, ProjectID string) *pubsub.Client {
	c, err := client.NewClient(ctx, ProjectID)
	if err != nil {
		fmt.Printf("unable to connect to server - %v", err)
		os.Exit(23)
	}

	return c


}
