package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type querySnippetsClient struct {
	*redash.Client
}

const listQuerySnippetsResBody = `[{"something":"something"}]`

func (c *querySnippetsClient) ListQuerySnippets(_ *redash.ListQuerySnippetsInput) *redash.ListQuerySnippetsOutput {
	return &redash.ListQuerySnippetsOutput{StatusCode: 200, Body: listQuerySnippetsResBody}
}

func TestNewCmdListquerySnippets(t *testing.T) {
	testClient := &querySnippetsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listQuerySnippetsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListQuerySnippets(testClient)
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

func (c *querySnippetsClient) GetQuerySnippet(_ *redash.GetQuerySnippetInput) *redash.GetQuerySnippetOutput {
	return &redash.GetQuerySnippetOutput{StatusCode: 200, Body: getQuerySnippetResBody}
}

func TestNewCmdGetQuerySnippet(t *testing.T) {
	testClient := &querySnippetsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--query-snippet-id", "1"}, want: getQuerySnippetResBody},
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
