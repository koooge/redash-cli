package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetDashboardList(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-dashboardlist",
		Short: "get-dashboardlist",
		Long:  `Get dashboardlist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetDashboardList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetDashboard(c redash.IClient) *cobra.Command {
	var dashboardSlug string

	var cmd = &cobra.Command{
		Use:   "get-dashboard",
		Short: "get-dashboard",
		Long:  `Get dashboard`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDashboardInput{
				DashboardSlug: dashboardSlug,
			}
			output := c.GetDashboard(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().StringVar(&dashboardSlug, "dashboard-slug", "", "dashboard slug")

	return cmd
}

func NewCmdGetPublicDashboard(c redash.IClient) *cobra.Command {
	var token string

	var cmd = &cobra.Command{
		Use:   "get-publicdashboard",
		Short: "get-publicdashboard",
		Long:  `Get publicdashboard`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetPublicDashboardInput{
				Token: token,
			}
			output := c.GetPublicDashboard(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().StringVar(&token, "token", "", "token")

	return cmd
}

func NewCmdGetDashboardTags(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-dashboardtags",
		Short: "get-dashboardtags",
		Long:  `Get dashboardtags`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetDashboardTags()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
