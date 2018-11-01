package cmd

import (
	"github.com/spf13/cobra"
)

func NewCmdQueryList() *cobra.Command {
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
