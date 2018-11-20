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

const listQueriesResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *queriesClient) ListQueries(_ *redash.ListQueriesInput) *redash.ListQueriesOutput {
	return &redash.ListQueriesOutput{StatusCode: 200, Body: listQueriesResBody}
}

func TestNewCmdListQueries(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listQueriesResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListQueries(testClient)
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

func (c *queriesClient) GetQuerySearch(_ *redash.GetQuerySearchInput) *redash.GetQuerySearchOutput {
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

func (c *queriesClient) GetQueryRecent(_ *redash.GetQueryRecentInput) *redash.GetQueryRecentOutput {
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

func (c *queriesClient) GetMyQueries(_ *redash.GetMyQueriesInput) *redash.GetMyQueriesOutput {
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

func (c *queriesClient) GetQueryTags(_ *redash.GetQueryTagsInput) *redash.GetQueryTagsOutput {
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

const createQueryResBody = `{"something":"something"}`

func (c *queriesClient) CreateQuery(_ *redash.CreateQueryInput) *redash.CreateQueryOutput {
	return &redash.CreateQueryOutput{StatusCode: 200, Body: createQueryResBody}
}

func TestNewCmdCreateQuery(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "1", "--query", "something", "--name", "something"}, want: createQueryResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdCreateQuery(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const modifyQueryResBody = `{"something":"something"}`

func (c *queriesClient) ModifyQuery(_ *redash.ModifyQueryInput) *redash.ModifyQueryOutput {
	return &redash.ModifyQueryOutput{StatusCode: 200, Body: modifyQueryResBody}
}

func TestNewCmdModifyQuery(t *testing.T) {
	testClient := &queriesClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--data-source-id", "1", "--query", "something", "--name", "something"}, want: modifyQueryResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdModifyQuery(testClient)
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
