package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type eventsClient struct {
	*redash.Client
}

const getEventsResBody = `[{"something":"something"}]`

func (c *eventsClient) GetEvents() *redash.GetEventsOutput {
	return &redash.GetEventsOutput{StatusCode: 200, Body: getEventsResBody}
}

func TestNewCmdGetEvents(t *testing.T) {
	testClient := &eventsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getEventsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetEvents(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
