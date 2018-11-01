package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetUserList() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-userlist",
		Short: "get-userlist",
		Long:  `Get userlist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := client.GetUserList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetUser() *cobra.Command {
	var userId int
	var cmd = &cobra.Command{
		Use:   "get-user",
		Short: "get-user",
		Long:  `Get user`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetUserInput{
				UserId: userId,
			}
			output := client.GetUser(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&userId, "user-id", 0, "user id")

	return cmd
}
