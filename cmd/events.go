package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdGetEvents() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-events",
		Short: "get-events",
		Long:  `Get events`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetEvents()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
