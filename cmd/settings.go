package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetOrganizationSettings(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-organization-settings",
		Short: "get-organization-settings",
		Long:  `Get OrganizationSettings`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetOrganizationSettingsInput{}
			output := c.GetOrganizationSettings(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
