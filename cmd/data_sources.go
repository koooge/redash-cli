package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListDataSources(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-data-sources",
		Short: "list-data-sources",
		Long:  `Get DataSourceList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListDataSourcesInput{}
			output := c.ListDataSources(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetDataSource(c redash.IClient) *cobra.Command {
	var dataSourceId int

	var cmd = &cobra.Command{
		Use:   "get-data-source",
		Short: "get-data-source",
		Long:  `Get DataSource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDataSourceInput{
				DataSourceId: dataSourceId,
			}
			output := c.GetDataSource(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "data-source-id", 0, "DataSource id")

	return cmd
}

func NewCmdGetDataSourceSchema(c redash.IClient) *cobra.Command {
	var dataSourceId int

	var cmd = &cobra.Command{
		Use:   "get-data-source-schema",
		Short: "get-data-source-schema",
		Long:  `Get DataSourceSchema`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDataSourceSchemaInput{
				DataSourceId: dataSourceId,
			}
			output := c.GetDataSourceSchema(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "data-source-id", 0, "DataSource id")

	return cmd
}
