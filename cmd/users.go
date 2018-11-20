package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListUsers(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-users",
		Short: "list-users",
		Long:  `Get UserList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListUsersInput{}
			output := c.ListUsers(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetUser(c redash.IClient) *cobra.Command {
	var userId int
	var cmd = &cobra.Command{
		Use:   "get-user",
		Short: "get-user",
		Long:  `Get User`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetUserInput{
				UserId: userId,
			}
			output := c.GetUser(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&userId, "user-id", 0, "User id")

	return cmd
}
