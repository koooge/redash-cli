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

	client = redash.NewClient(nil)

	// dashboards
	rootCmd.AddCommand(NewCmdGetDashboardList(client))
	rootCmd.AddCommand(NewCmdGetDashboard(client))
	rootCmd.AddCommand(NewCmdGetPublicDashboard(client))
	rootCmd.AddCommand(NewCmdGetDashboardTags(client))

	// data_sources
	rootCmd.AddCommand(NewCmdGetDataSourceList(client))
	rootCmd.AddCommand(NewCmdGetDataSource(client))
	rootCmd.AddCommand(NewCmdGetDataSourceSchema(client))

	// destinations
	rootCmd.AddCommand(NewCmdGetDestinationList(client))
	rootCmd.AddCommand(NewCmdGetDestination(client))
	rootCmd.AddCommand(NewCmdGetDestinationTypeList(client))

	// queries
	rootCmd.AddCommand(NewCmdGetQueryList(client))
	rootCmd.AddCommand(NewCmdGetQuery(client))
	rootCmd.AddCommand(NewCmdGetQuerySearch(client))
	rootCmd.AddCommand(NewCmdGetQueryRecent(client))
	rootCmd.AddCommand(NewCmdGetMyQueries(client))
	rootCmd.AddCommand(NewCmdGetQueryTags(client))
	rootCmd.AddCommand(NewCmdPostQueryList(client))
	rootCmd.AddCommand(NewCmdPostQuery(client))
	rootCmd.AddCommand(NewCmdDeleteQuery(client))

	// query_results
	rootCmd.AddCommand(NewCmdGetQueryResult(client))
	rootCmd.AddCommand(NewCmdGetJob(client))

	// query_snippets
	rootCmd.AddCommand(NewCmdGetQuerySnippetList(client))
	rootCmd.AddCommand(NewCmdGetQuerySnippet(client))

	// users
	rootCmd.AddCommand(NewCmdGetUserList(client))
	rootCmd.AddCommand(NewCmdGetUser(client))

	// groups
	rootCmd.AddCommand(NewCmdGetGroupList(client))
	rootCmd.AddCommand(NewCmdGetGroup(client))
	rootCmd.AddCommand(NewCmdGetGroupMemberList(client))

	// events
	rootCmd.AddCommand(NewCmdGetEvents(client))

	// alerts
	rootCmd.AddCommand(NewCmdGetAlertList(client))
	rootCmd.AddCommand(NewCmdGetAlert(client))
	rootCmd.AddCommand(NewCmdGetAlertSubscriptionList(client))

	// settings
	rootCmd.AddCommand(NewCmdGetOrganizationSettings(client))

	// favorites
	rootCmd.AddCommand(NewCmdGetQueryFavoriteList(client))
	rootCmd.AddCommand(NewCmdGetDashboardFavoriteList(client))

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
	client.SetConfig(config)
}
