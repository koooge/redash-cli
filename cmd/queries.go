package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetQueryList() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-querylist",
		Short: "get-querylist",
		Long:  `Get querylist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetQueryList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetQuery() *cobra.Command {
	var queryId int
	var cmd = &cobra.Command{
		Use:   "get-query",
		Short: "get-query",
		Long:  `Get query`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQueryInput{
				QueryId: queryId,
			}
			output := client.GetQuery(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&queryId, "query-id", 0, "query id")

	return cmd
}

func NewCmdGetQuerySearch() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-querysearch",
		Short: "get-querysearch",
		Long:  `Get querysearch`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetQuerySearch()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetQueryRecent() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-queryrecent",
		Short: "get-queryrecent",
		Long:  `Get queryrecent`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetQueryRecent()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetMyQueries() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-myqueries",
		Short: "get-myqueries",
		Long:  `Get myqueries`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetMyQueries()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}
