package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListQuerySnippets(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-query-snippets",
		Short: "list-query-snippets",
		Long:  `Get QuerySnippetList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListQuerySnippetsInput{}
			output := c.ListQuerySnippets(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetQuerySnippet(c redash.IClient) *cobra.Command {
	var querySnippetId int

	var cmd = &cobra.Command{
		Use:   "get-query-snippet",
		Short: "get-query-snippet",
		Long:  `Get QuerySnippet`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQuerySnippetInput{
				QuerySnippetId: querySnippetId,
			}
			output := c.GetQuerySnippet(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&querySnippetId, "query-snippet-id", 0, "QuerySnippet id")

	return cmd
}
