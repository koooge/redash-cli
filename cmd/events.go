package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetEvents(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-events",
		Short: "get-events",
		Long:  `Get Events`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetEventsInput{}
			output := c.GetEvents(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
