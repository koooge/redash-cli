package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListQueries(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-queries",
		Short: "list-queries",
		Long:  `Get QueryList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListQueriesInput{}
			output := c.ListQueries(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetQuery(c redash.IClient) *cobra.Command {
	var queryId int

	var cmd = &cobra.Command{
		Use:   "get-query",
		Short: "get-query",
		Long:  `Get Query`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQueryInput{
				QueryId: queryId,
			}
			output := c.GetQuery(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&queryId, "query-id", 0, "Query id")

	return cmd
}

func NewCmdGetQuerySearch(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-query-search",
		Short: "get-query-search",
		Long:  `Get QuerySearch`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQuerySearchInput{}
			output := c.GetQuerySearch(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetQueryRecent(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-query-recent",
		Short: "get-query-recent",
		Long:  `Get QueryRecent`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQueryRecentInput{}
			output := c.GetQueryRecent(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetMyQueries(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-my-queries",
		Short: "get-my-queries",
		Long:  `Get my-queries`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetMyQueriesInput{}
			output := c.GetMyQueries(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetQueryTags(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-query-tags",
		Short: "get-query-tags",
		Long:  `Get QueryTags`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetQueryTagsInput{}
			output := c.GetQueryTags(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdCreateQuery(c redash.IClient) *cobra.Command {
	var dataSourceId int
	var query string
	var name string
	var description string
	var schedule string

	var cmd = &cobra.Command{
		Use:   "create-query",
		Short: "create-query",
		Long:  `Post QueryList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.CreateQueryInput{
				DataSourceId: dataSourceId,
				Query:        query,
				Name:         name,
				Description:  description,
				Schedule:     schedule,
			}
			output := c.CreateQuery(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "data-source-id", 0, "DataSource id")
	cmd.Flags().StringVar(&query, "query", "", "Query")
	cmd.Flags().StringVar(&name, "name", "", "Name")
	cmd.Flags().StringVar(&description, "description", "", "Description")
	cmd.Flags().StringVar(&schedule, "schedule", "", "Schedule")

	return cmd
}

func NewCmdModifyQuery(c redash.IClient) *cobra.Command {
	var queryId int
	var dataSourceId int
	var query string
	var name string
	var description string
	var schedule string

	var cmd = &cobra.Command{
		Use:   "modify-query",
		Short: "modify-query",
		Long:  `Post Query`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ModifyQueryInput{
				QueryId:      queryId,
				DataSourceId: dataSourceId,
				Query:        query,
				Name:         name,
				Description:  description,
				Schedule:     schedule,
			}
			output := c.ModifyQuery(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&queryId, "query-id", 0, "Query id")
	cmd.Flags().IntVar(&dataSourceId, "data-source-id", 0, "DataSource id")
	cmd.Flags().StringVar(&query, "query", "", "Query")
	cmd.Flags().StringVar(&name, "name", "", "Name")
	cmd.Flags().StringVar(&description, "description", "", "Description")
	cmd.Flags().StringVar(&schedule, "schedule", "", "Schedule")

	return cmd
}

func NewCmdDeleteQuery(c redash.IClient) *cobra.Command {
	var queryId int

	var cmd = &cobra.Command{
		Use:   "delete-query",
		Short: "delete-query",
		Long:  `Delete query`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.DeleteQueryInput{
				QueryId: queryId,
			}
			output := c.DeleteQuery(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&queryId, "query-id", 0, "Query id")

	return cmd
}
