package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetDashboardList() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-dashboardlist",
		Short: "get-dashboardlist",
		Long:  `Get dashboardlist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetDashboardList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetDashboard() *cobra.Command {
	var dashboardSlug string

	var cmd = &cobra.Command{
		Use:   "get-dashboard",
		Short: "get-dashboard",
		Long:  `Get dashboard`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDashboardInput{
				DashboardSlug: dashboardSlug,
			}
			output := client.GetDashboard(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().StringVar(&dashboardSlug, "dashboard-slug", "", "dashboard slug")

	return cmd
}
