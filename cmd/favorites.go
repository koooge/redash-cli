package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListQueryFavorites(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-query-favorites",
		Short: "list-query-favorites",
		Long:  `Get QueryFavoriteList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListQueryFavoritesInput{}
			output := c.ListQueryFavorites(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdListDashboardFavorites(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-dashboard-favorites",
		Short: "list-dashboard-favorites",
		Long:  `Get DashboardFavoriteList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListDashboardFavoritesInput{}
			output := c.ListDashboardFavorites(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
