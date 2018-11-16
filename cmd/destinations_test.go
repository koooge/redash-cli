package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type destinationsClient struct {
	*redash.Client
}

const getDestinationListResBody = `[{"something":"something"}]`

func (c *destinationsClient) GetDestinationList() *redash.GetDestinationListOutput {
	return &redash.GetDestinationListOutput{StatusCode: 200, Body: getDestinationListResBody}
}

func TestNewCmdGetDestinationList(t *testing.T) {
	testClient := &destinationsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getDestinationListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDestinationList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getDestinationResBody = `{"something":"something"}`

func (c *destinationsClient) GetDestination(_ *redash.GetDestinationInput) *redash.GetDestinationOutput {
	return &redash.GetDestinationOutput{StatusCode: 200, Body: getDestinationResBody}
}

func TestNewCmdGetDestination(t *testing.T) {
	testClient := &destinationsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--destination-id", "1"}, want: getDestinationResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDestination(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getDestinationTypeListResBody = `[{"something":"something"}]`

func (c *destinationsClient) GetDestinationTypeList() *redash.GetDestinationTypeListOutput {
	return &redash.GetDestinationTypeListOutput{StatusCode: 200, Body: getDestinationTypeListResBody}
}

func TestNewCmdGetDestinationTypeList(t *testing.T) {
	testClient := &destinationsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getDestinationTypeListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDestinationTypeList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
