package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetEvents(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-events",
		Short: "get-events",
		Long:  `Get events`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetEvents()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
