package cmd

import (
	"bufio"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdConfig() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "config",
		Short: "config",
		Long:  `config`,
		RunE: func(cmd *cobra.Command, args []string) error {
			scanner := bufio.NewScanner(os.Stdin)

			cmd.Print("endpoint_url: ")
			scanner.Scan()
			endpointUrl := scanner.Text()
			if endpointUrl != "" {
				viper.Set(profile+".endpoint_url", endpointUrl)
			}

			cmd.Print("api_key: ")
			scanner.Scan()
			apiKey := scanner.Text()
			if apiKey != "" {
				viper.Set(profile+".api_key", apiKey)
			}

			err := updateConfig()
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}

func updateConfig() error {
	if err := os.MkdirAll(configPath, 0777); err != nil {
		return err
	}
	if err := viper.WriteConfig(); err != nil {
		return err
	}

	return nil
}
