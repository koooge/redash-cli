package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type queriesClient struct {
	*redash.Client
}

const getQueryListResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *queriesClient) GetQueryList() *redash.GetQueryListOutput {
	return &redash.GetQueryListOutput{StatusCode: 200, Body: getQueryListResBody}
}

func TestNewCmdGetQueryList(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getQueryListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQueryList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getQueryResBody = `{"something":"something"}`

func (c *queriesClient) GetQuery(_ *redash.GetQueryInput) *redash.GetQueryOutput {
	return &redash.GetQueryOutput{StatusCode: 200, Body: getQueryResBody}
}

func TestNewCmdGetQuery(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--query-id", "1"}, want: getQueryResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQuery(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getQuerySearchResBody = `[{"something":"something"}]`

func (c *queriesClient) GetQuerySearch() *redash.GetQuerySearchOutput {
	return &redash.GetQuerySearchOutput{StatusCode: 200, Body: getQuerySearchResBody}
}

func TestNewCmdGetQuerySearch(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getQuerySearchResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQuerySearch(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getQueryRecentResBody = `[{"something":"something"}]`

func (c *queriesClient) GetQueryRecent() *redash.GetQueryRecentOutput {
	return &redash.GetQueryRecentOutput{StatusCode: 200, Body: getQueryRecentResBody}
}

func TestNewCmdGetQueryRecent(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getQueryRecentResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQueryRecent(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getMyQueriesResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *queriesClient) GetMyQueries() *redash.GetMyQueriesOutput {
	return &redash.GetMyQueriesOutput{StatusCode: 200, Body: getMyQueriesResBody}
}

func TestNewCmdGetMyQueries(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getMyQueriesResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetMyQueries(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getQueryTagsResBody = `{"something":"something"}`

func (c *queriesClient) GetQueryTags() *redash.GetQueryTagsOutput {
	return &redash.GetQueryTagsOutput{StatusCode: 200, Body: getQueryTagsResBody}
}

func TestNewCmdGetQueryTags(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getQueryTagsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQueryTags(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const postQueryListResBody = `{"something":"something"}`

func (c *queriesClient) PostQueryList(_ *redash.PostQueryListInput) *redash.PostQueryListOutput {
	return &redash.PostQueryListOutput{StatusCode: 200, Body: postQueryListResBody}
}

func TestNewCmdPostQueryList(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--datasource-id", "1", "--query", "something", "--name", "something"}, want: postQueryListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdPostQueryList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const postQueryResBody = `{"something":"something"}`

func (c *queriesClient) PostQuery(_ *redash.PostQueryInput) *redash.PostQueryOutput {
	return &redash.PostQueryOutput{StatusCode: 200, Body: postQueryResBody}
}

func TestNewCmdPostQuery(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--datasource-id", "1", "--query", "something", "--name", "something"}, want: postQueryResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdPostQuery(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const deleteQueryResBody = `{"something":"something"}`

func (c *queriesClient) DeleteQuery(_ *redash.DeleteQueryInput) *redash.DeleteQueryOutput {
	return &redash.DeleteQueryOutput{StatusCode: 200, Body: deleteQueryResBody}
}

func TestNewCmdDeleteQuery(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--query-id", "1"}, want: deleteQueryResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdDeleteQuery(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
