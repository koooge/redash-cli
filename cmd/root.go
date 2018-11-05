package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/koooge/redash-sdk-go/redash"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command
var version string

var configFile string
var profile string
var endpointUrl string
var apiKey string
var client *redash.Client

var configPath string
var configFilename string

func NewCmdRoot() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "redash",
		Short: "redash",
		Long:  `redash`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default \"$HOME/.config/redash-cli/config.yaml\")")
	cmd.PersistentFlags().StringVar(&profile, "profile", "default", "profile")
	cmd.PersistentFlags().StringVar(&endpointUrl, "endpoint-url", "", "endpoint url")
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
	rootCmd.AddCommand(NewCmdVersion())
	rootCmd.AddCommand(NewCmdConfig())

	// dashboards
	rootCmd.AddCommand(NewCmdGetDashboardList())
	rootCmd.AddCommand(NewCmdGetDashboard())

	// data_sources
	rootCmd.AddCommand(NewCmdGetDataSourceList())
	rootCmd.AddCommand(NewCmdGetDataSource())
	rootCmd.AddCommand(NewCmdGetDataSourceSchema())

	// destinations
	rootCmd.AddCommand(NewCmdGetDestinationList())
	rootCmd.AddCommand(NewCmdGetDestination())
	rootCmd.AddCommand(NewCmdGetDestinationTypeList())

	// queries
	rootCmd.AddCommand(NewCmdGetQueryList())
	rootCmd.AddCommand(NewCmdGetQuery())
	rootCmd.AddCommand(NewCmdGetQuerySearch())
	rootCmd.AddCommand(NewCmdGetQueryRecent())
	rootCmd.AddCommand(NewCmdGetMyQueries())
	rootCmd.AddCommand(NewCmdGetQueryTags())
	rootCmd.AddCommand(NewCmdPostQueryList())
	rootCmd.AddCommand(NewCmdPostQuery())
	rootCmd.AddCommand(NewCmdDeleteQuery())

	// query_results
	rootCmd.AddCommand(NewCmdGetQueryResult())
	rootCmd.AddCommand(NewCmdGetJob())

	// query_snippets
	rootCmd.AddCommand(NewCmdGetQuerySnippetList())
	rootCmd.AddCommand(NewCmdGetQuerySnippet())

	// users
	rootCmd.AddCommand(NewCmdGetUserList())
	rootCmd.AddCommand(NewCmdGetUser())

	// groups
	rootCmd.AddCommand(NewCmdGetGroupList())
	rootCmd.AddCommand(NewCmdGetGroup())

	// events
	rootCmd.AddCommand(NewCmdGetEvents())

	// alerts
	rootCmd.AddCommand(NewCmdGetAlertList())
	rootCmd.AddCommand(NewCmdGetAlert())

	// settings
	rootCmd.AddCommand(NewCmdGetOrganizationSettings())

	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initClient)
}

func initConfig() {
	if configFile != "" {
		configPath = filepath.Dir(configFile)
		configFilename = filepath.Base(configFile)
		configFile = filepath.Join(configPath, configFilename)
		viper.SetConfigFile(configFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configPath = filepath.Join(home, ".config", "redash-cli")
		configFilename = "config.yaml"
		configFile = filepath.Join(configPath, configFilename)
		viper.SetConfigFile(configFile)
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()
}

func initClient() {
	if endpointUrl == "" {
		endpointUrl = viper.GetString(profile + ".endpoint_url")
	}
	if apiKey == "" {
		apiKey = viper.GetString(profile + ".api_key")
	}

	config := &redash.Config{
		EndpointUrl: endpointUrl,
		ApiKey:      apiKey,
	}
	client = redash.NewClient(config)
}
