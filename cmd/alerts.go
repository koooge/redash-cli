package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetAlertList(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-alertlist",
		Short: "get-alertlist",
		Long:  `Get alertlist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetAlertList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetAlert(c redash.IClient) *cobra.Command {
	var alertId int
	var cmd = &cobra.Command{
		Use:   "get-alert",
		Short: "get-alert",
		Long:  `Get alert`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetAlertInput{
				AlertId: alertId,
			}
			output := c.GetAlert(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&alertId, "alert-id", 0, "alert id")

	return cmd
}
