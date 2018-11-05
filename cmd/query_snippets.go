package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetQuerySnippetList() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-querysnippetlist",
		Short: "get-querysnippetlist",
		Long:  `Get querysnippetlist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetQuerySnippetList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetQuerySnippet() *cobra.Command {
	var querySnippetId int
	var cmd = &cobra.Command{
		Use:   "get-querysnippet",
		Short: "get-querysnippet",
		Long:  `Get querysnippet`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQuerySnippetInput{
				QuerySnippetId: querySnippetId,
			}
			output := client.GetQuerySnippet(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&querySnippetId, "querysnippet-id", 0, "querysnippet id")

	return cmd
}