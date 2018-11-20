package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListAlerts(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-alerts",
		Short: "list-alerts",
		Long:  `Get AlertList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListAlertsInput{}
			output := c.ListAlerts(input)
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
		Long:  `Get Alert`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetAlertInput{
				AlertId: alertId,
			}
			output := c.GetAlert(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&alertId, "alert-id", 0, "Alert id")

	return cmd
}

func NewCmdListAlertSubscriptions(c redash.IClient) *cobra.Command {
	var alertId int
	var cmd = &cobra.Command{
		Use:   "list-alert-subscriptions",
		Short: "list-alert-subscriptions",
		Long:  `Get AlertSubscriptionList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListAlertSubscriptionsInput{
				AlertId: alertId,
			}
			output := c.ListAlertSubscriptions(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&alertId, "alert-id", 0, "Alert id")

	return cmd
}
