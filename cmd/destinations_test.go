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

const listDestinationsResBody = `[{"something":"something"}]`

func (c *destinationsClient) ListDestinations(*redash.ListDestinationsInput) *redash.ListDestinationsOutput {
	return &redash.ListDestinationsOutput{StatusCode: 200, Body: listDestinationsResBody}
}

func TestNewCmdListDestinations(t *testing.T) {
	testClient := &destinationsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listDestinationsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListDestinations(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const createDestinationResBody = `{"something":"something"}`

func (c *destinationsClient) CreateDestination(*redash.CreateDestinationInput) *redash.CreateDestinationOutput {
	return &redash.CreateDestinationOutput{StatusCode: 200, Body: createDestinationResBody}
}

func TestNewCmdCreateDestination(t *testing.T) {
	testClient := &destinationsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--addresses", "foo@bar.piyo", "--name", "something", "--type", "email"}, want: createDestinationResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdCreateDestination(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const listDestinationsTypesResBody = `[{"something":"something"}]`

func (c *destinationsClient) ListDestinationsTypes(_ *redash.ListDestinationsTypesInput) *redash.ListDestinationsTypesOutput {
	return &redash.ListDestinationsTypesOutput{StatusCode: 200, Body: listDestinationsTypesResBody}
}

func TestNewCmdListDestinationsTypes(t *testing.T) {
	testClient := &destinationsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listDestinationsTypesResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListDestinationsTypes(testClient)
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

const updateDestinationResBody = `{"something":"something"}`

func (c *destinationsClient) UpdateDestination(_ *redash.UpdateDestinationInput) *redash.UpdateDestinationOutput {
	return &redash.UpdateDestinationOutput{StatusCode: 200, Body: updateDestinationResBody}
}

func TestNewCmdUpdateDestination(t *testing.T) {
	testClient := &destinationsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--destination-id", "1", "--addresses", "foo@bar.piyo", "--name", "something", "--type", "email"}, want: updateDestinationResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdUpdateDestination(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const deleteDestinationResBody = `{"something":"something"}`

func (c *destinationsClient) DeleteDestination(_ *redash.DeleteDestinationInput) *redash.DeleteDestinationOutput {
	return &redash.DeleteDestinationOutput{StatusCode: 200, Body: deleteDestinationResBody}
}

func TestNewCmdDeleteDestination(t *testing.T) {
	testClient := &destinationsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--destination-id", "1"}, want: deleteDestinationResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdDeleteDestination(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
