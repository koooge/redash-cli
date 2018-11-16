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
