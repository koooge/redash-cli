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

func NewCmdGetQueryTags() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-querytags",
		Short: "get-querytags",
		Long:  `Get querytags`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetQueryTags()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdPostQueryList() *cobra.Command {
	var dataSourceId int
	var query string
	var name string
	var description string
	var schedule string

	var cmd = &cobra.Command{
		Use:   "post-querylist",
		Short: "post-querylist",
		Long:  `Post querylist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.PostQueryListInput{
				DataSourceId: dataSourceId,
				Query:        query,
				Name:         name,
				Description:  description,
				Schedule:     schedule,
			}
			output := client.PostQueryList(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "datasource-id", 0, "datasource id")
	cmd.Flags().StringVar(&query, "query", "", "query")
	cmd.Flags().StringVar(&name, "name", "", "name")
	cmd.Flags().StringVar(&description, "description", "", "description")
	cmd.Flags().StringVar(&schedule, "schedule", "", "schedule")

	return cmd
}

func NewCmdPostQuery() *cobra.Command {
	var queryId int
	var dataSourceId int
	var query string
	var name string
	var description string
	var schedule string

	var cmd = &cobra.Command{
		Use:   "post-query",
		Short: "post-query",
		Long:  `Post query`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.PostQueryInput{
				QueryId:      queryId,
				DataSourceId: dataSourceId,
				Query:        query,
				Name:         name,
				Description:  description,
				Schedule:     schedule,
			}
			output := client.PostQuery(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&queryId, "query-id", 0, "query id")
	cmd.Flags().IntVar(&dataSourceId, "datasource-id", 0, "datasource id")
	cmd.Flags().StringVar(&query, "query", "", "query")
	cmd.Flags().StringVar(&name, "name", "", "name")
	cmd.Flags().StringVar(&description, "description", "", "description")
	cmd.Flags().StringVar(&schedule, "schedule", "", "schedule")

	return cmd
}

func NewCmdDeleteQuery() *cobra.Command {
	var queryId int

	var cmd = &cobra.Command{
		Use:   "delete-query",
		Short: "delete-query",
		Long:  `Delete query`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.DeleteQueryInput{
				QueryId: queryId,
			}
			output := client.DeleteQuery(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&queryId, "query-id", 0, "query id")

	return cmd
}
