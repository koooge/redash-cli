package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListDestinations(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-destinations",
		Short: "list-destinations",
		Long:  `ListDestinations`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListDestinationsInput{}
			output := c.ListDestinations(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdCreateDestination(c redash.IClient) *cobra.Command {
	var addresses string
	var name string
	var type_ string

	var cmd = &cobra.Command{
		Use:   "create-destination",
		Short: "create-destination",
		Long:  `CreateDestination`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.CreateDestinationInput{
				Options: &redash.CreateDestinationInputOptions{
					Addresses: addresses,
				},
				Name: name,
				Type: type_,
			}
			output := c.CreateDestination(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().StringVar(&addresses, "addresses", "", "Addresses")
	cmd.Flags().StringVar(&name, "name", "", "Name")
	cmd.Flags().StringVar(&type_, "type", "", "Type")

	return cmd
}

func NewCmdListDestinationsTypes(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-destinations-types",
		Short: "list-destinations-types",
		Long:  `ListDestinationsType`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListDestinationsTypesInput{}
			output := c.ListDestinationsTypes(input)
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
		Long:  `GetDestination`,
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

func NewCmdUpdateDestination(c redash.IClient) *cobra.Command {
	var destinationId int
	var addresses string
	var name string
	var type_ string

	var cmd = &cobra.Command{
		Use:   "update-destination",
		Short: "update-destination",
		Long:  `UpdateDestination`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.UpdateDestinationInput{
				DestinationId: destinationId,
				Options: &redash.UpdateDestinationInputOptions{
					Addresses: addresses,
				},
				Name: name,
				Type: type_,
			}
			output := c.UpdateDestination(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&destinationId, "destination-id", 0, "Destination id")
	cmd.Flags().StringVar(&addresses, "addresses", "", "Addresses")
	cmd.Flags().StringVar(&name, "name", "", "Name")
	cmd.Flags().StringVar(&type_, "type", "", "Type")

	return cmd
}

func NewCmdDeleteDestination(c redash.IClient) *cobra.Command {
	var destinationId int

	var cmd = &cobra.Command{
		Use:   "delete-destination",
		Short: "delete-destination",
		Long:  `DeleteDestination`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.DeleteDestinationInput{
				DestinationId: destinationId,
			}
			output := c.DeleteDestination(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&destinationId, "destination-id", 0, "Destination id")

	return cmd
}
