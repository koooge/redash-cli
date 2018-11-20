package cmd

import (
	"github.com/koooge/redash-sdk-go/redash"
	"github.com/spf13/cobra"
)

func NewCmdListGroups(c redash.IClient) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list-groups",
		Short: "list-groups",
		Long:  `Get GroupList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListGroupsInput{}
			output := c.ListGroups(input)
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

	cmd.Flags().IntVar(&groupId, "group-id", 0, "Group id")

	return cmd
}

func NewCmdListGroupMembers(c redash.IClient) *cobra.Command {
	var groupId int
	var cmd = &cobra.Command{
		Use:   "list-group-members",
		Short: "list-group-members",
		Long:  `Get GroupMemberList`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input := &redash.ListGroupMembersInput{
				GroupId: groupId,
			}
			output := c.ListGroupMembers(input)
			cmd.Println(output.Body)

			return nil
		},
	}

	cmd.Flags().IntVar(&groupId, "group-id", 0, "Group id")

	return cmd
}
