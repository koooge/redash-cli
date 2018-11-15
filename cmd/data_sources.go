package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetDataSourceList(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-datasourcelist",
		Short: "get-datasourcelist",
		Long:  `Get datasourcelist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetDataSourceList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetDataSource(c redash.IClient) *cobra.Command {
	var dataSourceId int

	var cmd = &cobra.Command{
		Use:   "get-datasource",
		Short: "get-datasource",
		Long:  `Get datasource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDataSourceInput{
				DataSourceId: dataSourceId,
			}
			output := c.GetDataSource(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "datasource-id", 0, "datasource id")

	return cmd
}

func NewCmdGetDataSourceSchema(c redash.IClient) *cobra.Command {
	var dataSourceId int

	var cmd = &cobra.Command{
		Use:   "get-datasourceschema",
		Short: "get-datasourceschema",
		Long:  `Get datasourceschema`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetDataSourceSchemaInput{
				DataSourceId: dataSourceId,
			}
			output := c.GetDataSourceSchema(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "datasource-id", 0, "datasource id")

	return cmd
}
