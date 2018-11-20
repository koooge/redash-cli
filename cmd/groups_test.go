package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type groupsClient struct {
	*redash.Client
}

const listGroupsResBody = `[{"something":"something"}]`

func (c *groupsClient) ListGroups(_ *redash.ListGroupsInput) *redash.ListGroupsOutput {
	return &redash.ListGroupsOutput{StatusCode: 200, Body: listGroupsResBody}
}

func TestNewCmdListGroups(t *testing.T) {
	testClient := &groupsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listGroupsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListGroups(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getGroupResBody = `{"something":"something"}`

func (c *groupsClient) GetGroup(_ *redash.GetGroupInput) *redash.GetGroupOutput {
	return &redash.GetGroupOutput{StatusCode: 200, Body: getGroupResBody}
}

func TestNewCmdGetGroup(t *testing.T) {
	testClient := &groupsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--group-id", "1"}, want: getGroupResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetGroup(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const listGroupMembersResBody = `[{"something":"something"}]`

func (c *groupsClient) ListGroupMembers(_ *redash.ListGroupMembersInput) *redash.ListGroupMembersOutput {
	return &redash.ListGroupMembersOutput{StatusCode: 200, Body: listGroupMembersResBody}
}

func TestNewCmdListGroupMembers(t *testing.T) {
	testClient := &groupsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--group-id", "1"}, want: listGroupMembersResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListGroupMembers(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
