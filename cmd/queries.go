package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdQueryList() *cobra.Command {
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

func NewCmdQuery() *cobra.Command {
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

	cmd.Flags().IntVar(&queryId, "queryId", 0, "query id")

	return cmd
}
