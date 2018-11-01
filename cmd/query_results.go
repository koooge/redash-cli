package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetQueryResult() *cobra.Command {
	var queryResultId int

	var cmd = &cobra.Command{
		Use:   "get-queryresult",
		Short: "get-queryresult",
		Long:  `Get queryresult`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQueryResultInput{
				QueryResultId: queryResultId,
			}
			output := client.GetQueryResult(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&queryResultId, "queryresult-id", 0, "queryresult id")

	return cmd
}
