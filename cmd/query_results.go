package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetQueryResult(c redash.IClient) *cobra.Command {
	var queryResultId int

	var cmd = &cobra.Command{
		Use:   "get-query-result",
		Short: "get-query-result",
		Long:  `Get QueryResult`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQueryResultInput{
				QueryResultId: queryResultId,
			}
			output := c.GetQueryResult(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&queryResultId, "query-result-id", 0, "QueryResult id")

	return cmd
}

func NewCmdGetJob(c redash.IClient) *cobra.Command {
	var jobId int

	var cmd = &cobra.Command{
		Use:   "get-job",
		Short: "get-job",
		Long:  `Get Job`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetJobInput{
				JobId: jobId,
			}
			output := c.GetJob(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&jobId, "job-id", 0, "Job id")

	return cmd
}
