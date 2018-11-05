package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdGetQueryFavoriteList() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-queryfavoritelist",
		Short: "get-queryfavoritelist",
		Long:  `Get queryfavoritelist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetQueryFavoriteList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetDashboardFavoriteList() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-dashboardfavoritelist",
		Short: "get-dashboardfavoritelist",
		Long:  `Get dashboardfavoritelist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetDashboardFavoriteList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
