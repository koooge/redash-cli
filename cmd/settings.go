package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdGetOrganizationSettings() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-organizationsettings",
		Short: "get-organizationsettings",
		Long:  `Get organizationsettings`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetOrganizationSettings()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
