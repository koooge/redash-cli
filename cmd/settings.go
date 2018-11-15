package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetOrganizationSettings(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-organizationsettings",
		Short: "get-organizationsettings",
		Long:  `Get organizationsettings`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetOrganizationSettings()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
