package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListDataSources(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-data-sources",
		Short: "list-data-sources",
		Long:  `ListDataSource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListDataSourcesInput{}
			output := c.ListDataSources(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdCreateDataSource(c redash.IClient) *cobra.Command {
	var dbname string
	var name string
	var type_ string

	var cmd = &cobra.Command{
		Use:   "create-data-source",
		Short: "create-data-source",
		Long:  `CreateDataSource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.CreateDataSourceInput{
				Options: &redash.CreateDataSourceInputOptions{
					Dbname: dbname,
				},
				Name: name,
				Type: type_,
			}
			output := c.CreateDataSource(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().StringVar(&dbname, "dbname", "", "Dbname")
	cmd.Flags().StringVar(&name, "name", "", "Name")
	cmd.Flags().StringVar(&type_, "type", "", "Type")

	return cmd
}

func NewCmdListDataSourcesTypes(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-data-sources-types",
		Short: "list-data-sources-types",
		Long:  `ListDataSourcesTypes`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.ListDataSourcesTypes(nil)
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
		Long:  `GetDataSource`,
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

func NewCmdUpdateDataSource(c redash.IClient) *cobra.Command {
	var dataSourceId int
	var dbname string
	var name string
	var type_ string

	var cmd = &cobra.Command{
		Use:   "update-data-source",
		Short: "update-data-source",
		Long:  `UpdateDataSource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.UpdateDataSourceInput{
				DataSourceId: dataSourceId,
				Options: &redash.UpdateDataSourceInputOptions{
					Dbname: dbname,
				},
				Name: name,
				Type: type_,
			}
			output := c.UpdateDataSource(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "data-source-id", 0, "DataSource id")
	cmd.Flags().StringVar(&dbname, "dbname", "", "Db name")
	cmd.Flags().StringVar(&name, "name", "", "Name")
	cmd.Flags().StringVar(&type_, "type", "", "Type")

	return cmd
}

func NewCmdDeleteDataSource(c redash.IClient) *cobra.Command {
	var dataSourceId int

	var cmd = &cobra.Command{
		Use:   "delete-data-source",
		Short: "delete-data-source",
		Long:  `DeleteDataSource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.DeleteDataSourceInput{
				DataSourceId: dataSourceId,
			}
			output := c.DeleteDataSource(input)
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

func NewCmdPauseDataSource(c redash.IClient) *cobra.Command {
	var dataSourceId int

	var cmd = &cobra.Command{
		Use:   "pause-data-source",
		Short: "pause-data-source",
		Long:  `PauseDataSource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.PauseDataSourceInput{
				DataSourceId: dataSourceId,
			}
			output := c.PauseDataSource(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "data-source-id", 0, "DataSource id")

	return cmd
}

func NewCmdUnpauseDataSource(c redash.IClient) *cobra.Command {
	var dataSourceId int

	var cmd = &cobra.Command{
		Use:   "unpause-data-source",
		Short: "unpause-data-source",
		Long:  `UnpauseDataSource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.UnpauseDataSourceInput{
				DataSourceId: dataSourceId,
			}
			output := c.UnpauseDataSource(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "data-source-id", 0, "DataSource id")

	return cmd
}

func NewCmdTestDataSource(c redash.IClient) *cobra.Command {
	var dataSourceId int

	var cmd = &cobra.Command{
		Use:   "test-data-source",
		Short: "test-data-source",
		Long:  `TestDataSource`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.TestDataSourceInput{
				DataSourceId: dataSourceId,
			}
			output := c.TestDataSource(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&dataSourceId, "data-source-id", 0, "DataSource id")

	return cmd
}
