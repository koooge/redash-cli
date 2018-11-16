package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type alertsClient struct {
	*redash.Client
}

const getAlertListResBody = `[{"something":"something"}]`

func (c *alertsClient) GetAlertList() *redash.GetAlertListOutput {
	return &redash.GetAlertListOutput{StatusCode: 200, Body: getAlertListResBody}
}

func TestNewCmdGetAlertList(t *testing.T) {
	testClient := &alertsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getAlertListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetAlertList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getAlertResBody = `{"something":"something"}`

func (c *alertsClient) GetAlert(_ *redash.GetAlertInput) *redash.GetAlertOutput {
	return &redash.GetAlertOutput{StatusCode: 200, Body: getAlertResBody}
}

func TestNewCmdGetAlert(t *testing.T) {
	testClient := &alertsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--alert-id", "1"}, want: getAlertResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetAlert(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getAlertSubscriptionListResBody = `[{"something":"something"}]`

func (c *alertsClient) GetAlertSubscriptionList(_ *redash.GetAlertSubscriptionListInput) *redash.GetAlertSubscriptionListOutput {
	return &redash.GetAlertSubscriptionListOutput{StatusCode: 200, Body: getAlertSubscriptionListResBody}
}

func TestNewCmdGetAlertSubscriptionList(t *testing.T) {
	testClient := &alertsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--alert-id", "1"}, want: getAlertSubscriptionListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetAlertSubscriptionList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
