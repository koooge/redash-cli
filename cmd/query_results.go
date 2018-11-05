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

func NewCmdGetJob() *cobra.Command {
	var jobId int

	var cmd = &cobra.Command{
		Use:   "get-job",
		Short: "get-job",
		Long:  `Get job`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetJobInput{
				JobId: jobId,
			}
			output := client.GetJob(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&jobId, "job-id", 0, "job id")

	return cmd
}
