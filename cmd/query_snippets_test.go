package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type querysnippetsClient struct {
	*redash.Client
}

const getQuerySnippetListResBody = `[{"something":"something"}]`

func (c *querysnippetsClient) GetQuerySnippetList() *redash.GetQuerySnippetListOutput {
	return &redash.GetQuerySnippetListOutput{StatusCode: 200, Body: getQuerySnippetListResBody}
}

func TestNewCmdGetQuerySnippetList(t *testing.T) {
	testClient := &querysnippetsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getQuerySnippetListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQuerySnippetList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getQuerySnippetResBody = `{"something":"something"}`

func (c *querysnippetsClient) GetQuerySnippet(_ *redash.GetQuerySnippetInput) *redash.GetQuerySnippetOutput {
	return &redash.GetQuerySnippetOutput{StatusCode: 200, Body: getQuerySnippetResBody}
}

func TestNewCmdGetQuerySnippet(t *testing.T) {
	testClient := &querysnippetsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--querysnippet-id", "1"}, want: getQuerySnippetResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQuerySnippet(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
