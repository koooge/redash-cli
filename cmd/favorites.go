package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetQueryFavoriteList(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-queryfavoritelist",
		Short: "get-queryfavoritelist",
		Long:  `Get queryfavoritelist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetQueryFavoriteList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetDashboardFavoriteList(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-dashboardfavoritelist",
		Short: "get-dashboardfavoritelist",
		Long:  `Get dashboardfavoritelist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetDashboardFavoriteList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
