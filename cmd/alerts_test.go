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

const listAlertsResBody = `[{"something":"something"}]`

func (c *alertsClient) ListAlerts(_ *redash.ListAlertsInput) *redash.ListAlertsOutput {
	return &redash.ListAlertsOutput{StatusCode: 200, Body: listAlertsResBody}
}

func TestNewCmdListAlerts(t *testing.T) {
	testClient := &alertsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: listAlertsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListAlerts(testClient)
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

const listAlertSubscriptionsResBody = `[{"something":"something"}]`

func (c *alertsClient) ListAlertSubscriptions(_ *redash.ListAlertSubscriptionsInput) *redash.ListAlertSubscriptionsOutput {
	return &redash.ListAlertSubscriptionsOutput{StatusCode: 200, Body: listAlertSubscriptionsResBody}
}

func TestNewCmdListAlertSubscriptions(t *testing.T) {
	testClient := &alertsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--alert-id", "1"}, want: listAlertSubscriptionsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdListAlertSubscriptions(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
