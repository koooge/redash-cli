package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type usersClient struct {
	*redash.Client
}

const listUsersResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *usersClient) ListUsers(_ *redash.ListUsersInput) *redash.ListUsersOutput {
	return &redash.ListUsersOutput{StatusCode: 200, Body: listUsersResBody}
}

func TestNewCmdListUsers(t *testing.T) {
	testClient := &usersClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listUsersResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListUsers(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getUserResBody = `{"something":"something"}`

func (c *usersClient) GetUser(_ *redash.GetUserInput) *redash.GetUserOutput {
	return &redash.GetUserOutput{StatusCode: 200, Body: getUserResBody}
}

func TestNewCmdGetUser(t *testing.T) {
	testClient := &usersClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--user-id", "1"}, want: getUserResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetUser(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
