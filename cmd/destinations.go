package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListDestinations(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-destinations",
		Short: "list-destinations",
		Long:  `Get DestinationList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListDestinationsInput{}
			output := c.ListDestinations(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetDestination(c redash.IClient) *cobra.Command {
	var destinationId int
	var cmd = &cobra.Command{
		Use:   "get-destination",
		Short: "get-destination",
		Long:  `Get Destination`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDestinationInput{
				DestinationId: destinationId,
			}
			output := c.GetDestination(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&destinationId, "destination-id", 0, "Destination id")

	return cmd
}

func NewCmdListDestinationsTypes(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-destinations-types",
		Short: "list-destinations-types",
		Long:  `Get DestinationsTypeList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListDestinationsTypesInput{}
			output := c.ListDestinationsTypes(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
