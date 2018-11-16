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

const getGroupListResBody = `[{"something":"something"}]`

func (c *groupsClient) GetGroupList() *redash.GetGroupListOutput {
	return &redash.GetGroupListOutput{StatusCode: 200, Body: getGroupListResBody}
}

func TestNewCmdGetGroupList(t *testing.T) {
	testClient := &groupsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getGroupListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetGroupList(testClient)
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

const getGroupMemberListResBody = `[{"something":"something"}]`

func (c *groupsClient) GetGroupMemberList(_ *redash.GetGroupMemberListInput) *redash.GetGroupMemberListOutput {
	return &redash.GetGroupMemberListOutput{StatusCode: 200, Body: getGroupMemberListResBody}
}

func TestNewCmdGetGroupMemberList(t *testing.T) {
	testClient := &groupsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--group-id", "1"}, want: getGroupMemberListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetGroupMemberList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
