package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type queryresultsClient struct {
	*redash.Client
}

const getQueryResultResBody = `{"something":"something"}`

func (c *queryresultsClient) GetQueryResult(_ *redash.GetQueryResultInput) *redash.GetQueryResultOutput {
	return &redash.GetQueryResultOutput{StatusCode: 200, Body: getQueryResultResBody}
}

func TestNewCmdGetQueryResult(t *testing.T) {
	testClient := &queryresultsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--queryresult-id", "1"}, want: getQueryResultResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetQueryResult(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getJobResBody = `{"something":"something"}`

func (c *queryresultsClient) GetJob(_ *redash.GetJobInput) *redash.GetJobOutput {
	return &redash.GetJobOutput{StatusCode: 200, Body: getJobResBody}
}

func TestNewCmdGetJob(t *testing.T) {
	testClient := &queryresultsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--job-id", "1"}, want: getJobResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetJob(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
