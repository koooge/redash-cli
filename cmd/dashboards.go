package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListDashboards(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-dashboards",
		Short: "list-dashboards",
		Long:  `Get DashboardList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListDashboardsInput{}
			output := c.ListDashboards(input)
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
		Long:  `Get Dashboard`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDashboardInput{
				DashboardSlug: dashboardSlug,
			}
			output := c.GetDashboard(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().StringVar(&dashboardSlug, "dashboard-slug", "", "Dashboard slug")

	return cmd
}

func NewCmdGetPublicDashboard(c redash.IClient) *cobra.Command {
	var token string

	var cmd = &cobra.Command{
		Use:   "get-public-dashboard",
		Short: "get-public-dashboard",
		Long:  `Get PublicDashboard`,
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
		Use:   "get-dashboard-tags",
		Short: "get-dashboard-tags",
		Long:  `Get DashboardTags`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDashboardTagsInput{}
			output := c.GetDashboardTags(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
