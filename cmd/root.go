package cmd

import (
	"fmt"
	"os"

	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

var endpointUrl string
var apiKey string
var client *redash.Client

func NewCmdRoot() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "redash",
		Short: "redash",
		Long:  `redash`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVar(&endpointUrl, "endpoint-url", "http://localhost", "endpoint url")
	cmd.PersistentFlags().StringVar(&apiKey, "apikey", "", "api key")

	return cmd
}

func Execute() {
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd = NewCmdRoot()

	rootCmd.AddCommand(NewCmdQueryList())

	cobra.OnInitialize(initClient)
}

func initClient() {
	config := &redash.Config{
		EndpointUrl: endpointUrl,
		ApiKey:      apiKey,
	}
	client = redash.NewClient(config)
}
