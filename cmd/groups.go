package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdGetGroupList(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-grouplist",
		Short: "get-grouplist",
		Long:  `Get grouplist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			output := c.GetGroupList()
			cmd.Println(output.Body)

			return nil
		},
	}

	return cmd
}

func NewCmdGetGroup(c redash.IClient) *cobra.Command {
	var groupId int
	var cmd = &cobra.Command{
		Use:   "get-group",
		Short: "get-group",
		Long:  `Get group`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetGroupInput{
				GroupId: groupId,
			}
			output := c.GetGroup(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&groupId, "group-id", 0, "group id")

	return cmd
}

func NewCmdGetGroupMemberList(c redash.IClient) *cobra.Command {
	var groupId int
	var cmd = &cobra.Command{
		Use:   "get-groupmemberlist",
		Short: "get-groupmemberlist",
		Long:  `Get groupmemberlist`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.GetGroupMemberListInput{
				GroupId: groupId,
			}
			output := c.GetGroupMemberList(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&groupId, "group-id", 0, "group id")

	return cmd
}
