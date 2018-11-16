package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/koooge/redash-sdk-go/redash"
)

type dashboardsClient struct {
	*redash.Client
}

const getDashboardListResBody = `{"something":"something","results":[{"something":"something"}]}`

func (c *dashboardsClient) GetDashboardList() *redash.GetDashboardListOutput {
	return &redash.GetDashboardListOutput{StatusCode: 200, Body: getDashboardListResBody}
}

func TestNewCmdGetDashboardList(t *testing.T) {
	testClient := &dashboardsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getDashboardListResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDashboardList(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getDashboardResBody = `{"something":"something"}`

func (c *dashboardsClient) GetDashboard(_ *redash.GetDashboardInput) *redash.GetDashboardOutput {
	return &redash.GetDashboardOutput{StatusCode: 200, Body: getDashboardResBody}
}

func TestNewCmdGetDashboard(t *testing.T) {
	testClient := &dashboardsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--dashboard-slug", "something"}, want: getDashboardResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDashboard(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getPublicDashboardResBody = `{"something":"something"}`

func (c *dashboardsClient) GetPublicDashboard(_ *redash.GetPublicDashboardInput) *redash.GetPublicDashboardOutput {
	return &redash.GetPublicDashboardOutput{StatusCode: 200, Body: getPublicDashboardResBody}
}

func TestNewCmdGetPublicDashboard(t *testing.T) {
	testClient := &dashboardsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{"--token", "something"}, want: getPublicDashboardResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetPublicDashboard(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}

const getDashboardTagsResBody = `{"something":"something"}`

func (c *dashboardsClient) GetDashboardTags() *redash.GetDashboardTagsOutput {
	return &redash.GetDashboardTagsOutput{StatusCode: 200, Body: getDashboardTagsResBody}
}

func TestNewCmdGetDashboardTags(t *testing.T) {
	testClient := &dashboardsClient{}

	testCases := []struct {
		args []string
		want string
	}{
		{args: []string{}, want: getDashboardTagsResBody},
	}

	for _, c := range testCases {
		buf := new(bytes.Buffer)
		cmd := NewCmdGetDashboardTags(testClient)
		cmd.SetOutput(buf)
		cmd.SetArgs(c.args)
		cmd.Execute()

		get := buf.String()
		if c.want != strings.TrimRight(get, "\n") {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
		}
	}
}
