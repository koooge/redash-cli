package cmd

import "github.com/spf13/cobra"

func NewCmdVersion() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "version",
		Short: "version",
		Long:  `version`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println(version)
			return nil
		},
	}

	return cmd
}
