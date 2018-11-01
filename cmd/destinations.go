package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetDestinationList() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-destinationlist",
		Short: "get-destinationlist",
		Long:  `Get destinationlist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetDestinationList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetDestination() *cobra.Command {
	var destinationId int
	var cmd = &cobra.Command{
		Use:   "get-destination",
		Short: "get-destination",
		Long:  `Get destination`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDestinationInput{
				DestinationId: destinationId,
			}
			output := client.GetDestination(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&destinationId, "destination-id", 0, "destination id")

	return cmd
}
